package http

import (
	"Task-Service/pkg/domain"
	"Task-Service/pkg/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type TaskEndpoints struct {
	TaskService *service.TaskService
}


var (
	ErrTokenNotFound   = errors.New("no se encontró el token")
	ErrTokenInvalid    = errors.New("error en el token")
	ErrUserIDNotFound  = errors.New("no se encontró user_id en el token")
)


var jwtKey = []byte("tok3n2024") // La clave debe coincidir con la que usas para firmar el token

func extractUserID(c *gin.Context) (uint, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, ErrTokenNotFound
	}

	tokenString = tokenString[len("Bearer "):]

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, ErrTokenInvalid
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, ErrUserIDNotFound
	}

	return uint(userIDFloat), nil
}



func (e *TaskEndpoints) CreateTaskHandler(c *gin.Context) {
	userID, err := extractUserID(c)
	if err != nil {
		switch err {
		case ErrTokenNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token en la cabecera"})
			return
		case ErrTokenInvalid:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		case ErrUserIDNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No se encontró información de usuario en el token"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inesperado"})
			return
		}
	}

	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	task.UserID = userID
	if err := e.TaskService.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (e *TaskEndpoints) GetTaskAllHandler(c *gin.Context) {
	userID, err := extractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks, err := e.TaskService.GetTaskAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}


func (e *TaskEndpoints) DeleteTask(c *gin.Context) {
	userID, err := extractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	// Obtener el taskID de la URL
	taskID := c.Param("taskID")

	// Verificar que el taskID no esté vacío
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "taskID es obligatorio"})
		return
	}

	// Convertir taskID a uint
	taskIDUint, err := strconv.ParseUint(taskID, 10, 32) // uint de 32 bits
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "taskID no es válido"})
		return
	}

	// Convertir a uint
	taskIDUint32 := uint(taskIDUint)

	// Llamar al servicio para eliminar la tarea
	err = e.TaskService.DeleteTask(taskIDUint32, userID)
	if err != nil {
		// Si ocurre un error en el servicio, puedes devolver el error adecuado
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la tarea"})
		return
	}

	// Si la tarea se elimina con éxito
	c.JSON(http.StatusOK, gin.H{"message": "La tarea fue eliminada exitosamente"})
}
