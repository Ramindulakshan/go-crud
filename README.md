# go-crud

Bootstrap for the User Management REST API described in `SIMPLE_USER_MANAGEMENT_FLOW.md`.

## What is included

- Go API skeleton with `Handler -> Service -> Repository` flow
- User CRUD endpoints (`/users`)
- In-memory repository for quick start (DB repository can be added next)
- PostgreSQL Docker Compose and SQL migration scaffold
- OpenAPI starter file in `docs/openapi.yaml`

## Run locally

1. Copy env file:

```bash
cp .env.example .env
```

2. Start PostgreSQL (optional for now, since repo is in-memory):

```bash
docker compose up -d
```

3. Run API:

```bash
go mod tidy
go run ./cmd/server
```

API will start at `http://localhost:8080`.

## Endpoints

- `GET /health`
- `POST /users`
- `GET /users`
- `GET /users/{id}`
- `PUT /users/{id}`
- `DELETE /users/{id}`
