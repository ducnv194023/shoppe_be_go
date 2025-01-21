-- name: GetUserProfileByUserID :one
SELECT * FROM user_profiles
WHERE user_id = ? AND deleted_at IS NULL;

-- name: GetUserProfileByUUID :one
SELECT * FROM user_profiles
WHERE uuid = ? AND deleted_at IS NULL;

-- name: CreateUserProfile :execresult
INSERT INTO user_profiles (
    uuid, user_id, full_name, avatar_url, gender, 
    date_of_birth, bio, country_code, language, timezone,
    social_links, preferences, metadata
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateUserProfile :exec
UPDATE user_profiles
SET 
    full_name = COALESCE(?, full_name),
    avatar_url = COALESCE(?, avatar_url),
    gender = COALESCE(?, gender),
    date_of_birth = COALESCE(?, date_of_birth),
    bio = COALESCE(?, bio),
    country_code = COALESCE(?, country_code),
    language = COALESCE(?, language),
    timezone = COALESCE(?, timezone),
    social_links = COALESCE(?, social_links),
    preferences = COALESCE(?, preferences),
    metadata = COALESCE(?, metadata)
WHERE user_id = ? AND deleted_at IS NULL;

-- name: DeleteUserProfile :exec
UPDATE user_profiles
SET deleted_at = CURRENT_TIMESTAMP
WHERE user_id = ? AND deleted_at IS NULL;