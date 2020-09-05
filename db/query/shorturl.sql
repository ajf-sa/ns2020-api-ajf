-- name: CreateShortUrl :exec
INSERT INTO
    shorturl (origin, short)
VALUES
    ($1,$2);

-- name: GetShortUrl :one
SELECT * FROM shorturl
WHERE id = $1 LIMIT 1;


-- name: ListShortUrl :many
SELECT * FROM shorturl
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateShortUrl :exec
UPDATE shorturl SET hits = $2
WHERE id = $1;


-- name: DeleteShortUrl :exec
DELETE FROM shorturl WHERE id = $1;