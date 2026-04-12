package handlers

import (
	"gin-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req struct {
		Name        string `json:"name"`
		Descritpion string `json:"description"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	task, err := h.service.CreateTask(c, req.Name, req.Descritpion)

	if err != nil {
		c.JSON(500, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}
