# Go Server Development Tracking

## Timeline & Progress
- **2025-11-05** – Completed development plan outlining architecture, data store, and concurrency considerations.
- **2025-11-06** – Stubbed Go project structure, initialized `go.mod`, and verified local build with `go run main.go`.
- **2025-11-07** – Implemented CRUD handlers aligned with shared REST contract; added in-memory store guarded by `sync.RWMutex`.
- **2025-11-08** – Created unit tests for handlers covering success/error cases; fixed JSON tag mismatch discovered during testing.
- **2025-11-09** – Validated server with client tester suite; documented setup/run instructions in `README.md`.

## Key Decisions
- **Data Store**: Chose in-memory map keyed by UUID to simplify state management during assignment; persistence out of scope.
- **Concurrency Control**: Added `sync.RWMutex` around shared store to ensure thread safety without excessive locking.
- **Routing**: Leveraged standard library `net/http` and `http.ServeMux` to avoid external dependencies; sufficient for defined endpoints.
- **Error Handling**: Centralized JSON error responses to maintain consistent structure across handlers.
- **Testing Strategy**: Adopted table-driven tests to cover permutations of valid/invalid payloads and missing IDs.

## Outstanding Items
- Consider adding persistence abstraction for future extensions (e.g., file or DB backend).
- Investigate adding structured logging for easier debugging during integration tests.
- Monitor Go module updates; rerun `go mod tidy` before final submission to lock dependencies.

## Code Walk Summary
- **Date**: 2025-11-10
- **Participants**: Implementer – Samuel Coe; Reviewer – Runyuan Feng
- **Implementer Platform**: macOS Sonoma
- **Reviewer Platform**: Windows 11
- **Reviewer Verification**: Server started successfully via `go run main.go` and passed CRUD regression tests on reviewer machine.
- **Key Feedback**:
  - Document the mutex rationale in code comments for future maintainers.
  - Add structured logging around request entry/exit to aid debugging.
  - Consider extracting validation logic into a separate helper to simplify handlers.

