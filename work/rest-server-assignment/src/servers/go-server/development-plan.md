- **Approach**
  - Reuse the shared REST contract from the TypeScript reference server to ensure consistent routes and payloads.
  - Model users as an in-memory store using a map keyed by UUID strings to keep implementation simple and stateless.
  - Implement request handlers with idiomatic Go `net/http`, using helper functions for parsing JSON and sending responses.
  - Add middleware-like helpers for validating IDs and request bodies to reduce duplication.
  - Write table-driven unit tests for the handler functions covering happy paths and error cases.

- **Key Components**
  - `main.go`: entry point; configures router, registers handlers, and starts the HTTP server on port 5003.
  - `handlers.go`: defines handlers for CRUD operations (`GET /users`, `GET /users/{id}`, `POST`, `PUT`, `PATCH`, `DELETE`).
  - `models.go`: contains `User` struct definition and helper methods for validation.
  - `store.go`: abstracts data layer (in-memory map) with functions to list, create, update, and delete users.
  - `errors.go`: centralizes error response formatting and status code mapping.
  - `router.go`: optional helper to compose routes; uses standard library multiplexer.

- **Language-Specific Considerations**
  - Use Go modules (`go.mod`) to manage dependencies; keep third-party libraries minimal to simplify setup.
  - Handle JSON encoding/decoding with Go’s `encoding/json`, ensuring struct tags align with API schema.
  - Carefully manage pointer vs value semantics when updating the in-memory store to avoid unintended copies.
  - Ensure all handlers are concurrent-safe; protect shared store with a `sync.RWMutex`.
  - Provide clear error messages and HTTP status codes leveraging Go error wrapping patterns.
  - Include formatting (`gofmt`) and linting (`golangci-lint`) in development workflow to maintain code quality.

