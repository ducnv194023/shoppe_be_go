-- name: GetTokenByID :one
SELECT * FROM user_tokens 
WHERE id = ? AND deleted_at IS NULL;

-- name: GetTokenByUUID :one
SELECT * FROM user_tokens 
WHERE uuid = ? AND deleted_at IS NULL;

-- name: GetTokenByToken :one
SELECT * FROM user_tokens 
WHERE token = ? AND deleted_at IS NULL;

-- name: ListUserTokens :many
SELECT * FROM user_tokens 
WHERE user_id = ? 
AND token_type = ?
AND is_revoked = false 
AND expires_at > NOW() 
AND deleted_at IS NULL;

-- name: CreateToken :execresult
INSERT INTO user_tokens (
    uuid, user_id, token_type, token,
    expires_at, ip_address, user_agent
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: RevokeToken :exec
UPDATE user_tokens 
SET 
    is_revoked = true,
    revoked_at = CURRENT_TIMESTAMP
WHERE token = ? 
AND deleted_at IS NULL;

-- name: RevokeAllUserTokens :exec
UPDATE user_tokens 
SET 
    is_revoked = true,
    revoked_at = CURRENT_TIMESTAMP
WHERE user_id = ? 
AND token_type = ?
AND is_revoked = false 
AND deleted_at IS NULL;

-- name: DeleteExpiredTokens :exec
DELETE FROM user_tokens 
WHERE expires_at < NOW() 
OR is_revoked = true;

-- name: ValidateToken :one
SELECT * FROM user_tokens 
WHERE token = ? 
AND token_type = ?
AND is_revoked = false 
AND expires_at > NOW() 
AND deleted_at IS NULL;

-- name: CountActiveTokens :one
SELECT COUNT(*) FROM user_tokens 
WHERE user_id = ? 
AND token_type = ?
AND is_revoked = false 
AND expires_at > NOW() 
AND deleted_at IS NULL;

-- name: DeleteToken :exec
UPDATE user_tokens
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;