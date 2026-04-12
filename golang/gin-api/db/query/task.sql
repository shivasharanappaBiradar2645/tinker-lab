-- name: CreateTask :one
INSERT INTO tasks (name,descritpion,status)
VALUES ($1,$2,$3)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;


-- name: ListTasks :many
SELECT * FROM tasks
ORDER by id;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;