package repository

import (
	"context"
	"database/sql"
	"gin-api/db/sqlc"
)

type TaskRepository struct {
	q *sqlc.Queries
}

func NewTaskRepository(q *sqlc.Queries) *TaskRepository {
	return &TaskRepository{
		q: q,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, name string, description string) (sqlc.Task, error) {
	task, err := r.q.CreateTask(ctx, sqlc.CreateTaskParams{
		Name: name,
		Descritpion: sql.NullString{
			String: description,
			Valid:  description != "",
		},
		Status: "INCOMPLETE",
	})
	return task, err
}

func (r *TaskRepository) GetTask(ctx context.Context, id int32) sqlc.Task {
	task, _ := r.q.GetTask(ctx, id)
	return task
}

func (r *TaskRepository) ListTasks(ctx context.Context) ([]sqlc.Task, error) {
	return r.q.ListTasks(ctx)
}

func (r *TaskRepository) DeletTask(ctx context.Context, id int32) error {
	return r.q.DeleteTask(ctx, id)
}
