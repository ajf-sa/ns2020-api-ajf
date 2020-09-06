-- name: CreateUser :exec
INSERT INTO
    users (username, password)
VALUES
    ($1,$2);

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;


-- name: ListUser :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateUser :exec
UPDATE users SET is_active = $2
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;