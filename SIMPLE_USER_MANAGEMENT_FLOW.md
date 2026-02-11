# User Management REST API -- Simple Flow Document

## Project Title

User Management REST API Development

## Objective

Develop a REST API for user management using Go.
The API handles CRUD operations with PostgreSQL integration and Swagger documentation.

------------------------------------------------------------------------

# Architecture Flow

Request -> Handler -> Service -> Repository(sqlc) -> PostgreSQL -> back

------------------------------------------------------------------------

# 1. Technical Stack

- Programming Language: Go
- Framework: Chi Router
- Database: PostgreSQL (Docker)
- Query Layer: sqlc
- Validation Library: go-playground/validator
- API Documentation: OpenAPI + Swagger UI
- Testing: go test
- Linting: golangci-lint

------------------------------------------------------------------------

# 2. Project Requirements

## 2.1 CRUD API for User Entity

Create a REST API that allows:

- Create User
- Read User
- Update User
- Delete User

### User Fields

- userId (UUID)
  - Primary Key
  - Auto-generated
- firstName (String)
  - Required
  - Min 2, Max 50 characters
- lastName (String)
  - Required
  - Min 2, Max 50 characters
- email (String)
  - Required
  - Valid Email Format
- phone (String)
  - Optional
  - Valid E.164 Phone Number (example: +14155552671)
- age (Integer)
  - Optional
  - Positive Integer
- status (Enum: Active, Inactive)
  - Optional
  - Default: Active

------------------------------------------------------------------------

## 2.2 API Endpoints

- POST /users -> Create user
- GET /users/{id} -> Retrieve user
- GET /users -> List all users
- PATCH /users/{id} -> Update user
- DELETE /users/{id} -> Delete user

------------------------------------------------------------------------

## 2.3 Database Configuration

- Use Docker to run PostgreSQL service.
- Use `.env` values for DB connection.
- Apply migration on startup (users table creation).

------------------------------------------------------------------------

## 2.4 Input Validation

- Validate request payload.
- Return appropriate error responses for invalid inputs.

------------------------------------------------------------------------

## 2.5 OpenAPI + Swagger

- OpenAPI spec file: `docs/openapi.yaml`
- Swagger UI endpoint: `/doc/`
- Raw OpenAPI spec endpoint: `/doc/openapi.yaml`

------------------------------------------------------------------------

## 2.6 sqlc Integration

- sqlc config file: `sqlc.yaml`
- SQL query source: `internal/db/query/users.sql`
- Generated package path: `internal/db/sqlc`
- Repository layer uses sqlc queries for all CRUD operations.

------------------------------------------------------------------------

## 2.7 Response Codes

- Use proper HTTP response codes:
  - 200 OK
  - 201 Created
  - 204 No Content
  - 400 Bad Request
  - 404 Not Found
  - 500 Internal Server Error

------------------------------------------------------------------------

## 2.8 Testing

- Unit tests for service and handler
- Integration tests for DB and API endpoints

------------------------------------------------------------------------

## 2.9 Linting & Code Quality

- Use golangci-lint
- Maintain clean project structure

------------------------------------------------------------------------

# Deliverables

1. Source Code Repository
2. Docker Setup for PostgreSQL
3. OpenAPI + Swagger endpoint
4. sqlc query + generated package integration
5. Unit and Integration Tests

------------------------------------------------------------------------

# Layer Responsibility (Request -> DB)

## Request

Client sends HTTP request.

## Handler

- Receives request
- Validates/parses input
- Calls service
- Sends HTTP response

## Service

- Contains business logic
- Calls repository layer

## Repository

- Uses sqlc-generated queries
- Interacts with PostgreSQL
- Maps DB models to API models

## DB

- PostgreSQL stores user data

## Back Flow

DB -> Repository -> Service -> Handler -> Response

------------------------------------------------------------------------

# Local Run Steps

1. `cp .env.example .env`
2. `docker compose up -d`
3. `go mod tidy`
4. `go run ./cmd/server`

Swagger:
- `http://localhost:8080/doc/`

------------------------------------------------------------------------

# Test Steps

1. Run unit tests:
   - `go test ./internal/user -v`
2. Run integration tests (with PostgreSQL running):
   - `TEST_DATABASE_URL="host=localhost port=5432 user=postgres password=postgres dbname=usersdb sslmode=disable" go test ./internal/http -run Integration -v`
