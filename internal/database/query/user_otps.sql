-- name: GetOTPByID :one
SELECT * FROM user_otps 
WHERE id = ? AND deleted_at IS NULL;

-- name: GetOTPByUUID :one
SELECT * FROM user_otps 
WHERE uuid = ? AND deleted_at IS NULL;

-- name: GetValidOTP :one
SELECT * FROM user_otps 
WHERE user_id = ? 
AND otp_type = ? 
AND otp_code = ? 
AND target = ?
AND is_used = false 
AND attempt_count < max_attempts 
AND expires_at > NOW() 
AND deleted_at IS NULL;

-- name: CreateOTP :execresult
INSERT INTO user_otps (
    uuid, user_id, otp_code, otp_type,
    target, expires_at, max_attempts
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: IncrementOTPAttempt :exec
UPDATE user_otps 
SET attempt_count = attempt_count + 1 
WHERE id = ? 
AND deleted_at IS NULL;

-- name: MarkOTPAsUsed :exec
UPDATE user_otps 
SET 
    is_used = true,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? 
AND deleted_at IS NULL;

-- name: InvalidateOldOTPs :exec
UPDATE user_otps 
SET 
    is_used = true,
    updated_at = CURRENT_TIMESTAMP
WHERE user_id = ? 
AND otp_type = ? 
AND target = ? 
AND is_used = false 
AND deleted_at IS NULL;

-- name: GetLatestOTP :one
SELECT * FROM user_otps 
WHERE user_id = ? 
AND otp_type = ? 
AND target = ? 
AND deleted_at IS NULL 
ORDER BY created_at DESC 
LIMIT 1;

-- name: CountRecentOTPs :one
SELECT COUNT(*) FROM user_otps 
WHERE user_id = ? 
AND otp_type = ? 
AND target = ? 
AND created_at > DATE_SUB(NOW(), INTERVAL ? MINUTE) 
AND deleted_at IS NULL;

-- name: DeleteExpiredOTPs :exec
DELETE FROM user_otps 
WHERE expires_at < NOW() 
OR (is_used = true AND created_at < DATE_SUB(NOW(), INTERVAL 24 HOUR));

-- name: DeleteOTP :exec
UPDATE user_otps
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;