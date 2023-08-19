-- name: GetUsers :many
SELECT
    id,
    username,
    email,
    avatar_url,
    oauth_id,
    oauth_type
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetUser :one
SELECT
    id,
    username,
    email,
    avatar_url,
    oauth_id,
    oauth_type
FROM users 
WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    hash_password,
    avatar_url,
    oauth_id,
    oauth_type
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING id;

-- name: UpdateUser :exec
UPDATE users SET
    username = $1,
    avatar_url = $2
WHERE id = $3 AND deleted_at IS NULL;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE deleted_at IS NOT NULL;