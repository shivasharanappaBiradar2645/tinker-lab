package routes

import (
	"gin-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, taskHandler *handlers.TaskHandler) {
	api := r.Group("/api")
	{
		api.GET("/tasks", taskHandler.CreateTask)
	}
}
