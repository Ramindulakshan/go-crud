-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  email,
  phone,
  age,
  status
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING user_id, first_name, last_name, email, phone, age, status, created_at, updated_at;

-- name: GetUserByID :one
SELECT user_id, first_name, last_name, email, phone, age, status, created_at, updated_at
FROM users
WHERE user_id = $1;

-- name: ListUsers :many
SELECT user_id, first_name, last_name, email, phone, age, status, created_at, updated_at
FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2,
    last_name = $3,
    email = $4,
    phone = $5,
    age = $6,
    status = $7,
    updated_at = NOW()
WHERE user_id = $1
RETURNING user_id, first_name, last_name, email, phone, age, status, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;
