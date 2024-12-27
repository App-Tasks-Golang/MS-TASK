package main

import (
	"Task-Service/internal/config"
	"Task-Service/pkg/adapters"
	"Task-Service/pkg/service"
	"Task-Service/transport/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	DB, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("error al obtener la conexión a la base de datos: %v", err)
	}

	taskRepo := adapters.NewTaskRepo(DB)
	taskService := service.NewTaskService(taskRepo)

	http.SetupRouter(r, taskService)


	r.Run(":8080")
}