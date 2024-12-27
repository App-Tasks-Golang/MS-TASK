package http

import (
	"Task-Service/pkg/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, taskService *service.TaskService) {

	// Crear una instancia de los endpoints
	endpoints := &TaskEndpoints{
		TaskService: taskService,
	}

	// Definir las rutas
	api := router.Group("/tasks")
	{
		api.POST("/", endpoints.CreateTaskHandler)
		api.GET("/", endpoints.GetTaskAllHandler)
		api.DELETE("/:taskID", endpoints.DeleteTask)
	}

}
