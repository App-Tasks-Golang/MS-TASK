package service

import (
	"Task-Service/pkg/domain"
)


// TaskService para manejar lógica de negocio
type TaskService struct {
	TaskRepo domain.TaskRepository
}

func NewTaskService(taskRepo domain.TaskRepository) *TaskService {
	return &TaskService{TaskRepo: taskRepo}
}

// Crear una tarea
func (s *TaskService) CreateTask(task domain.Task) error {
	return s.TaskRepo.CreateTask(task)
}

// Obtener todas las tareas de un usuario
func (s *TaskService) GetTaskAll(userID uint) ([]*domain.GetTasks, error) {
	return s.TaskRepo.GetTaskAll(userID)
}


// Actualizar una tarea específica
func (s *TaskService) UpdateTask(id uint, userID uint, updateTask domain.Task) error {
	return s.TaskRepo.UpdateTask(id, userID, updateTask)
}

// Eliminar una tarea
func (s *TaskService) DeleteTask(taskID uint, userID uint) error {
	return s.TaskRepo.DeleteTask(taskID, userID)
}
