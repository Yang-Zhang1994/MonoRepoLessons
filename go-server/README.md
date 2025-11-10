# Go REST Server

## Overview
This Go server implements the same RESTful API as the provided TypeScript and Python servers.  
It supports CRUD operations for users and passes the test suite in the client tester.

## Installation

### Requirements
- Go 1.20 or higher

```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go server.go
```

## Port
Runs on port 5003 by default.

## Routes Implemented

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | Get all users |
| GET | `/users/:id` | Get user by ID |
| POST | `/users` | Add a new user |
| PUT | `/users/:id` | Update user by ID |
| PATCH | `/users/:id/hours` | Add hours to a user |
| DELETE | `/users` | Delete all users |
| DELETE | `/users/:id` | Delete a specific user |

## Cross-Platform Test
- **Implementer**: macOS (Yang Zhang(Samuel))
- **Reviewer**: Windows (Runyuan Feng)
- **Status**: Passed all CRUD tests