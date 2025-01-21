-- name: GetUserAddressByID :one
SELECT * FROM user_addresses
WHERE id = ? AND deleted_at IS NULL;

-- name: GetUserAddressByUUID :one
SELECT * FROM user_addresses
WHERE uuid = ? AND deleted_at IS NULL;

-- name: ListUserAddresses :many
SELECT * FROM user_addresses
WHERE user_id = ? AND deleted_at IS NULL
ORDER BY is_default DESC, id DESC;

-- name: GetDefaultUserAddress :one
SELECT * FROM user_addresses
WHERE user_id = ? AND is_default = true AND deleted_at IS NULL;

-- name: CreateUserAddress :execresult
INSERT INTO user_addresses (
    uuid, user_id, receiver_name, receiver_phone,
    province_code, district_code, ward_code, street_address,
    label, is_default, notes
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateUserAddress :exec
UPDATE user_addresses
SET 
    receiver_name = COALESCE(?, receiver_name),
    receiver_phone = COALESCE(?, receiver_phone),
    province_code = COALESCE(?, province_code),
    district_code = COALESCE(?, district_code),
    ward_code = COALESCE(?, ward_code),
    street_address = COALESCE(?, street_address),
    label = COALESCE(?, label),
    is_default = COALESCE(?, is_default),
    notes = COALESCE(?, notes)
WHERE id = ? AND deleted_at IS NULL;

-- name: DeleteUserAddress :exec
UPDATE user_addresses
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;