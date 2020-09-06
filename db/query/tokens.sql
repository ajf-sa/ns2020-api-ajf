-- name: CreateToken :exec
INSERT INTO
    tokens (key, user_id)
VALUES
    ($1,$2);

-- name: GetToken :one
SELECT * FROM tokens
WHERE id = $1 LIMIT 1;


-- name: ListToken :many
SELECT * FROM tokens
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateToken :exec
UPDATE tokens SET is_active = $2
WHERE id = $1;


-- name: DeleteToken :exec
DELETE FROM tokens WHERE id = $1;