-- name: GetUserBasicByID :one
SELECT * FROM user_basics
WHERE id = ? AND deleted_at IS NULL;

-- name: GetUserBasicByUUID :one
SELECT * FROM user_basics
WHERE uuid = ? AND deleted_at IS NULL;

-- name: GetUserBasicByEmail :one
SELECT * FROM user_basics
WHERE email = ? AND deleted_at IS NULL;

-- name: GetUserBasicByPhone :one
SELECT * FROM user_basics
WHERE phone = ? AND deleted_at IS NULL;

-- name: ListUserBasics :many
SELECT * FROM user_basics
WHERE deleted_at IS NULL
ORDER BY id
LIMIT ? OFFSET ?;

-- name: CreateUserBasic :execresult
INSERT INTO user_basics (
    uuid, email, phone, password_hash, role, status
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: UpdateUserBasic :exec
UPDATE user_basics
SET 
    email = COALESCE(?, email),
    phone = COALESCE(?, phone),
    password_hash = COALESCE(?, password_hash),
    role = COALESCE(?, role),
    status = COALESCE(?, status)
WHERE id = ? AND deleted_at IS NULL;

-- name: VerifyUserEmail :exec
UPDATE user_basics
SET email_verified_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;

-- name: VerifyUserPhone :exec
UPDATE user_basics
SET phone_verified_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteUserBasic :exec
UPDATE user_basics
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;