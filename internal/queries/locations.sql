-- name: GetLocations :many
SELECT id, address, latitude, longitude, category, description, created_at, updated_at
FROM locations;


-- name: GetLocationsByCategory :many
SELECT id, address, latitude, longitude, category, description, created_at, updated_at
FROM locations
WHERE category LIKE ?
ORDER BY created_at DESC;


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
