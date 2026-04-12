package main

import (
	"database/sql"
	"gin-api/db/sqlc"
	"gin-api/handlers"
	"gin-api/repository"
	"gin-api/routes"
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	constring := "sfdfs"
	db, _ := sql.Open("postgres", constring)

	queries := sqlc.Queries(db)
	repo := repository.NewTaskRepository(queries)
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(*service)

	r := gin.Default()

	r.POST("/tasks", handler.CreateTask)
	routes.SetupRoutes(r, handler)
}
