-- name: GetLocations :many
SELECT  id, address, latitude, longitude, category, description, created_at, updated_at
FROM locations;


-- name: AddLocations :one
INSERT INTO locations
    (address, latitude, longitude, category, description)
VALUES (?, ?, ?, ?, ?)
RETURNING *;


-- name: UpdateLocation :one
UPDATE locations
SET latitude = ?, longitude = ?, description = ?, category = ?
WHERE id = ?
RETURNING *;
