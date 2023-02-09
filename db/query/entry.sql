-- name: CreateEntry :one
INSERT INTO entries
    (account_id, amount)
VALUES
    ($1, $2)
RETURNING *;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;