-- name: GetRooms :many
SELECT
    rooms.id,
    room_id,
    owner_id,
    is_private,
    max_players
    
FROM rooms
LEFT JOIN users ON rooms.owner_id = users.id
WHERE rooms.deleted_at IS NULL AND rooms.expired_at < NOW()
ORDER BY rooms.created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetRoom :one
SELECT
    id,
    room_id,
    owner_id,
    is_private,
    max_players
FROM rooms
WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateRoom :one
INSERT INTO rooms (
    room_id,
    owner_id,
    is_private,
    max_players
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING id;

-- name: DeleteRooms :exec
DELETE FROM rooms
WHERE deleted_at IS NOT NULL;