package adapters

import (
	"Task-Service/pkg/domain"
	"gorm.io/gorm"
)

// Repositorio
type TaskRepo struct {
	DB *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{DB: db}
}

// Crear tarea
func (r *TaskRepo) CreateTask(task domain.Task) error {
	return r.DB.Create(&task).Error
}

// Obtener todas las tareas de un usuario específico
func (r *TaskRepo) GetTaskAll(userID uint) ([]*domain.GetTasks, error) {
	var tasks []*domain.Task
	err := r.DB.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	// Mapear de Task a GetTasks
	var result []*domain.GetTasks
	for _, task := range tasks {
		result = append(result, &domain.GetTasks{
			IdTask:      task.ID,
			Title:       task.Title,
			Description: task.Description,
		})
	}

	return result, nil
}

// Actualizar una tarea
func (r *TaskRepo) UpdateTask(id uint, userID uint, updateTask domain.Task) error {
	var task domain.Task
	if err := r.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		return err
	}
	return r.DB.Model(&task).Updates(updateTask).Error
}

// Eliminar una tarea
func (r *TaskRepo) DeleteTask(taskID uint, userID uint) error {
	return r.DB.Where("id = ? AND user_id = ?", taskID, userID).Delete(&domain.Task{}).Error
}
