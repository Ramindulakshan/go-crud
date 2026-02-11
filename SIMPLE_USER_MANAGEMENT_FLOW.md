# User Management REST API -- Simple Flow Document

## Project Title

User Management REST API Development

## Objective

Develop a REST API for user management using Go programming language.\
The API should handle CRUD operations with DB integration and proper
documentation.

------------------------------------------------------------------------

# Architecture Flow

Request → Handler → Service → Repository → DB → back

------------------------------------------------------------------------

# 1. Technical Stack

-   Programming Language: Go
-   Framework: Chi Router (for handling REST API requests)
-   Database: PostgreSQL (running in Docker)
-   ORM/Query Builder: sqlc
-   Validation Library: go-playground validator
-   API Documentation: OpenAPI
-   Testing: go testing
-   Linting: golangci-lint

------------------------------------------------------------------------

# 2. Project Requirements

## 2.1 CRUD API for User Entity

Create a REST API that allows:

-   Create User
-   Read User
-   Update User
-   Delete User

### User Fields

-   userId (UUID)
    -   Primary Key\
    -   Auto-generated
-   firstName (String)
    -   Required\
    -   Min 2, Max 50 characters
-   lastName (String)
    -   Required\
    -   Min 2, Max 50 characters
-   email (String)
    -   Required\
    -   Valid Email Format
-   phone (String)
    -   Optional\
    -   Valid Phone Number
-   age (Integer)
    -   Optional\
    -   Positive Integer
-   status (Enum: Active, Inactive)
    -   Optional\
    -   Default: Active

------------------------------------------------------------------------

## 2.2 API Endpoints

-   POST /users → Create user
-   GET /users/{id} → Retrieve user
-   GET /users → List all users
-   PUT /users/{id} → Update user
-   DELETE /users/{id} → Delete user

------------------------------------------------------------------------

## 2.3 Database Configuration

-   Use Docker to run PostgreSQL service.
-   Use .env file to provide database connection details.

------------------------------------------------------------------------

## 2.4 Input Validation

-   Validate request payload.
-   Return appropriate error responses for invalid inputs.

------------------------------------------------------------------------

## 2.5 OpenAPI Documentation

-   Generate OpenAPI documentation for the REST API.
-   Provide documentation for each endpoint.

------------------------------------------------------------------------

## 2.6 Response Codes

-   Use proper HTTP response codes:
    -   200 OK
    -   201 Created
    -   400 Bad Request
    -   404 Not Found
    -   500 Internal Server Error

------------------------------------------------------------------------

## 2.7 Testing

-   Implement unit tests for service and handler functions.
-   Implement integration tests for database and API endpoints.

------------------------------------------------------------------------

## 2.8 Linting & Code Quality

-   Use golangci-lint for code linting and formatting.
-   Ensure clean code structure.

------------------------------------------------------------------------

# Deliverables

1.  Source Code Repository (GitHub or GitLab).
2.  Docker Setup (Docker Compose for PostgreSQL).
3.  API Documentation (OpenAPI).
4.  Unit & Integration Tests.
5.  Code Quality Report (if required).

------------------------------------------------------------------------

# Layer Responsibility (Request → DB)

## Request

Client sends HTTP request to API endpoint.

## Handler

-   Receives request.
-   Validates input.
-   Calls Service layer.
-   Sends response back.

## Service

-   Contains business logic.
-   Calls Repository layer.
-   Processes data before returning.

## Repository

-   Handles database queries using sqlc.
-   Interacts with PostgreSQL.
-   Returns data to Service.

## DB

-   PostgreSQL database.
-   Stores user data.

## Back Flow

DB → Repository → Service → Handler → Response to Client
