# go-crud

User Management REST API in Go with layered architecture:
`Handler -> Service -> Repository -> PostgreSQL`.

## What is included

- Chi-based REST API
- PostgreSQL-backed repository using `sqlc`-style generated queries (`internal/db/sqlc`)
- SQL query files for sqlc in `internal/db/query/users.sql`
- `sqlc` config in `sqlc.yaml`
- Startup migration for users table
- OpenAPI spec in `docs/openapi.yaml`
- Swagger UI at `/doc`

## Run locally

1. Copy env file:

```bash
cp .env.example .env
```

2. Start PostgreSQL:

```bash
docker compose up -d
```

3. Install dependencies and run API:

```bash
go mod tidy
go run ./cmd/server
```

API base URL: `http://localhost:8080`

## Swagger

- UI: `http://localhost:8080/doc/`
- Spec: `http://localhost:8080/doc/openapi.yaml`

## sqlc notes

- SQL source: `internal/db/query/users.sql`
- Config: `sqlc.yaml`
- Generated output target: `internal/db/sqlc`

If sqlc is installed in your environment, regenerate with:

```bash
sqlc generate
```

## Endpoints

- `GET /health`
- `POST /users`
- `GET /users`
- `GET /users/{id}`
- `PATCH /users/{id}`
- `DELETE /users/{id}`

## Testing

- Unit tests:

```bash
go test ./internal/user -v
```

- Integration tests (requires reachable PostgreSQL):

```bash
TEST_DATABASE_URL="host=localhost port=5432 user=postgres password=postgres dbname=usersdb sslmode=disable" go test ./internal/http -run Integration -v
```
