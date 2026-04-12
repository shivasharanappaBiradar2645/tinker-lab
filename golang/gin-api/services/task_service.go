package services

import (
	"context"
	"gin-api/db/sqlc"
	"gin-api/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (h *TaskService) CreateTask(ctx context.Context, name string, description string) (sqlc.Task, error) {
	return h.repo.CreateTask(ctx, name, description)
}
