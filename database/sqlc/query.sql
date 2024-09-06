-- name: GetImage :one
SELECT * FROM image
WHERE id = ? LIMIT 1;

-- name: ListImage :many
SELECT * FROM image
ORDER BY id;

-- name: CreateImage :one
INSERT INTO image (
  created_at, url
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateImage :exec
UPDATE image
set url = ?
WHERE id = ?;

-- name: DeleteImage :exec
DELETE FROM image
WHERE id = ?;

-- name: GetContainer :one
SELECT * FROM container
WHERE id = ? LIMIT 1;

-- name: ListContainer :many
SELECT * FROM container
ORDER BY name;

-- name: CreateContainer :one
INSERT INTO container (
  created_at, image_id, name
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateContainer :exec
UPDATE container
set name = ?,
image_id = ?
WHERE id = ?;

-- name: DeleteContainer :exec
DELETE FROM container
WHERE id = ?;
