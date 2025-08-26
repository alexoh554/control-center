-- name: CreateTask :one
INSERT INTO tasks (title, description, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET 
    title = COALESCE($1, title),
    description = COALESCE($2, description),
    status = COALESCE($3, status)
WHERE id = $4
RETURNING *;

-- name: DeleteTask :one
UPDATE tasks
SET deleted_at=CURRENT_TIMESTAMP
WHERE id=$1
RETURNING *;