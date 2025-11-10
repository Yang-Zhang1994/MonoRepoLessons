package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(handler http.Handler, method, target string, body []byte) *httptest.ResponseRecorder {
	var reader *bytes.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	} else {
		reader = bytes.NewReader([]byte{})
	}

	req := httptest.NewRequest(method, target, reader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func TestGetUsersReturnsEmptySlice(t *testing.T) {
	handler := NewServer()

	rr := performRequest(handler, http.MethodGet, "/users", nil)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var users []User
	if err := json.Unmarshal(rr.Body.Bytes(), &users); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(users) != 0 {
		t.Fatalf("expected empty list, got %v", users)
	}
}

func TestCreateAndFetchUser(t *testing.T) {
	handler := NewServer()

	createBody := []byte(`{"name":"Alice"}`)
	createRes := performRequest(handler, http.MethodPost, "/users", createBody)
	if createRes.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", createRes.Code)
	}

	var created User
	if err := json.Unmarshal(createRes.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to unmarshal create response: %v", err)
	}

	if created.ID != 1 || created.Name != "Alice" || created.HoursWorked != 0 {
		t.Fatalf("unexpected created user: %+v", created)
	}

	getRes := performRequest(handler, http.MethodGet, "/users/1", nil)
	if getRes.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", getRes.Code)
	}

	var fetched User
	if err := json.Unmarshal(getRes.Body.Bytes(), &fetched); err != nil {
		t.Fatalf("failed to unmarshal get response: %v", err)
	}

	if fetched != created {
		t.Fatalf("expected fetched user to equal created user, got %+v", fetched)
	}
}

func TestUpdateUserName(t *testing.T) {
	handler := NewServer()

	performRequest(handler, http.MethodPost, "/users", []byte(`{"name":"Alice"}`))

	updateBody := []byte(`{"name":"Alice Updated"}`)
	updateRes := performRequest(handler, http.MethodPut, "/users/1", updateBody)
	if updateRes.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", updateRes.Code)
	}

	var updated User
	if err := json.Unmarshal(updateRes.Body.Bytes(), &updated); err != nil {
		t.Fatalf("failed to unmarshal update response: %v", err)
	}

	if updated.Name != "Alice Updated" {
		t.Fatalf("expected name to be updated, got %s", updated.Name)
	}
}

func TestUpdateUserHours(t *testing.T) {
	handler := NewServer()

	performRequest(handler, http.MethodPost, "/users", []byte(`{"name":"Alice"}`))

	patchBody := []byte(`{"hoursToAdd":5}`)
	patchRes := performRequest(handler, http.MethodPatch, "/users/1", patchBody)
	if patchRes.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", patchRes.Code)
	}

	var patched User
	if err := json.Unmarshal(patchRes.Body.Bytes(), &patched); err != nil {
		t.Fatalf("failed to unmarshal patch response: %v", err)
	}

	if patched.HoursWorked != 5 {
		t.Fatalf("expected hoursWorked to be 5, got %d", patched.HoursWorked)
	}
}

func TestDeleteUser(t *testing.T) {
	handler := NewServer()

	performRequest(handler, http.MethodPost, "/users", []byte(`{"name":"Alice"}`))

	deleteRes := performRequest(handler, http.MethodDelete, "/users/1", nil)
	if deleteRes.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", deleteRes.Code)
	}

	var deleted User
	if err := json.Unmarshal(deleteRes.Body.Bytes(), &deleted); err != nil {
		t.Fatalf("failed to unmarshal delete response: %v", err)
	}

	if deleted.ID != 1 {
		t.Fatalf("expected deleted user ID 1, got %d", deleted.ID)
	}

	verifyRes := performRequest(handler, http.MethodGet, "/users/1", nil)
	if verifyRes.Code != http.StatusNotFound {
		t.Fatalf("expected status 404 after deletion, got %d", verifyRes.Code)
	}
}

func TestDeleteAllUsers(t *testing.T) {
	handler := NewServer()

	performRequest(handler, http.MethodPost, "/users", []byte(`{"name":"Alice"}`))
	performRequest(handler, http.MethodPost, "/users", []byte(`{"name":"Bob"}`))

	res := performRequest(handler, http.MethodDelete, "/users", nil)
	if res.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.Code)
	}

	var remaining []User
	if err := json.Unmarshal(res.Body.Bytes(), &remaining); err != nil {
		t.Fatalf("failed to unmarshal delete-all response: %v", err)
	}

	if len(remaining) != 0 {
		t.Fatalf("expected no users after delete all, got %v", remaining)
	}
}

