-- name: CreateTodo :one
INSERT INTO todos (title, status)
VALUES ($1, $2)
RETURNING *;
