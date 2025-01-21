-- name: GetUserPaymentMethodByID :one
SELECT * FROM user_payment_methods
WHERE id = ? AND deleted_at IS NULL;

-- name: ListUserPaymentMethods :many
SELECT * FROM user_payment_methods
WHERE user_id = ? AND deleted_at IS NULL
ORDER BY is_default DESC, id DESC;

-- name: GetDefaultUserPaymentMethod :one
SELECT * FROM user_payment_methods
WHERE user_id = ? AND is_default = true AND deleted_at IS NULL;

-- name: CreateUserPaymentMethod :execresult
INSERT INTO user_payment_methods (
    user_id, payment_type, provider, is_default,
    status, card_info, bank_info, wallet_info
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateUserPaymentMethod :exec
UPDATE user_payment_methods
SET 
    payment_type = COALESCE(?, payment_type),
    provider = COALESCE(?, provider),
    is_default = COALESCE(?, is_default),
    status = COALESCE(?, status),
    card_info = COALESCE(?, card_info),
    bank_info = COALESCE(?, bank_info),
    wallet_info = COALESCE(?, wallet_info)
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteUserPaymentMethod :exec
UPDATE user_payment_methods
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;