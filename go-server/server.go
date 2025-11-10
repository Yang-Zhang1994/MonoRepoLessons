package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type server struct {
	store *userStore
}

type userStore struct {
	mu     sync.RWMutex
	users  []User
	nextID int
}

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	HoursWorked int    `json:"hoursWorked"`
}

func NewServer() http.Handler {
	s := &server{
		store: &userStore{
			users:  make([]User, 0),
			nextID: 1,
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/users", s.handleUsers)
	mux.HandleFunc("/users/", s.handleUserByID)

	return s.withCORS(mux)
}

func (s *server) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.respondWithUsers(w)
	case http.MethodPost:
		s.createUser(w, r)
	case http.MethodDelete:
		s.deleteAllUsers(w)
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
	default:
		s.methodNotAllowed(w, "GET, POST, DELETE, OPTIONS")
	}
}

func (s *server) handleUserByID(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(segments) < 2 || segments[0] != "users" {
		s.notFound(w)
		return
	}

	id, err := strconv.Atoi(segments[1])
	if err != nil || id <= 0 {
		s.notFound(w)
		return
	}

	// Support PATCH /users/:id and PATCH /users/:id/hours
	if len(segments) == 3 {
		if segments[2] == "hours" && r.Method == http.MethodPatch {
			s.updateUserHours(w, r, id)
			return
		}
		s.notFound(w)
		return
	}

	if len(segments) > 2 {
		s.notFound(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.respondWithUser(w, id)
	case http.MethodPut:
		s.updateUser(w, r, id)
	case http.MethodPatch:
		s.updateUserHours(w, r, id)
	case http.MethodDelete:
		s.deleteUser(w, id)
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
	default:
		s.methodNotAllowed(w, "GET, PUT, PATCH, DELETE, OPTIONS")
	}
}

func (s *server) respondWithUsers(w http.ResponseWriter) {
	s.store.mu.RLock()
	defer s.store.mu.RUnlock()

	s.writeJSON(w, http.StatusOK, s.store.users)
}

func (s *server) respondWithUser(w http.ResponseWriter, id int) {
	s.store.mu.RLock()
	defer s.store.mu.RUnlock()

	for _, u := range s.store.users {
		if u.ID == id {
			s.writeJSON(w, http.StatusOK, u)
			return
		}
	}

	s.notFound(w)
}

func (s *server) createUser(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string `json:"name"`
	}

	if err := decodeJSON(r, &payload); err != nil {
		s.badRequest(w, err.Error())
		return
	}

	name := strings.TrimSpace(payload.Name)
	if name == "" {
		s.badRequest(w, "Name is required and must be a non-empty string")
		return
	}

	s.store.mu.Lock()
	user := User{
		ID:          s.store.nextID,
		Name:        name,
		HoursWorked: 0,
	}
	s.store.nextID++
	s.store.users = append(s.store.users, user)
	s.store.mu.Unlock()

	s.writeJSON(w, http.StatusCreated, user)
}

func (s *server) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var payload struct {
		Name *string `json:"name"`
	}

	if err := decodeJSON(r, &payload); err != nil {
		s.badRequest(w, err.Error())
		return
	}

	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	for i := range s.store.users {
		if s.store.users[i].ID == id {
			if payload.Name != nil {
				name := strings.TrimSpace(*payload.Name)
				if name == "" {
					s.badRequest(w, "Name is required and must be a non-empty string")
					return
				}
				s.store.users[i].Name = name
			}
			s.writeJSON(w, http.StatusOK, s.store.users[i])
			return
		}
	}

	s.notFound(w)
}

func (s *server) updateUserHours(w http.ResponseWriter, r *http.Request, id int) {
	var payload struct {
		HoursToAdd *int `json:"hoursToAdd"`
	}

	if err := decodeJSON(r, &payload); err != nil {
		s.badRequest(w, err.Error())
		return
	}

	if payload.HoursToAdd == nil {
		s.badRequest(w, "Invalid hoursToAdd value")
		return
	}

	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	for i := range s.store.users {
		if s.store.users[i].ID == id {
			s.store.users[i].HoursWorked += *payload.HoursToAdd
			s.writeJSON(w, http.StatusOK, s.store.users[i])
			return
		}
	}

	s.notFound(w)
}

func (s *server) deleteUser(w http.ResponseWriter, id int) {
	s.store.mu.Lock()
	defer s.store.mu.Unlock()

	for i := range s.store.users {
		if s.store.users[i].ID == id {
			deleted := s.store.users[i]
			s.store.users = append(s.store.users[:i], s.store.users[i+1:]...)
			s.writeJSON(w, http.StatusOK, deleted)
			return
		}
	}

	s.notFound(w)
}

func (s *server) deleteAllUsers(w http.ResponseWriter) {
	s.store.mu.Lock()
	s.store.users = make([]User, 0)
	s.store.nextID = 1
	s.store.mu.Unlock()

	s.writeJSON(w, http.StatusOK, []User{})
}

func decodeJSON(r *http.Request, dst any) error {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("Invalid JSON payload")
		}
		return errors.New("Invalid JSON payload")
	}

	if decoder.More() {
		return errors.New("Invalid JSON payload")
	}

	return nil
}

func (s *server) withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *server) writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if payload == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		// If writing fails, fall back to a generic error.
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *server) notFound(w http.ResponseWriter) {
	s.writeText(w, http.StatusNotFound, "User not found")
}

func (s *server) badRequest(w http.ResponseWriter, message string) {
	s.writeText(w, http.StatusBadRequest, message)
}

func (s *server) methodNotAllowed(w http.ResponseWriter, allow string) {
	w.Header().Set("Allow", allow)
	s.writeText(w, http.StatusMethodNotAllowed, "Method not allowed")
}

func (s *server) writeText(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(message))
}

