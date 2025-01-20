-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (
    username,
    email,
    password_hash
) VALUES (
    ?, ?, ?
);

-- name: UpdateUser :exec
UPDATE users 
SET 
    username = ?,
    email = ?,
    password_hash = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: ListUsersByCreatedAt :many
SELECT * FROM users
WHERE created_at >= ?
ORDER BY created_at DESC;