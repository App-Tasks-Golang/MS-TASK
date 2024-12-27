package domain


type TaskRepository interface {
	CreateTask(task Task) error
	GetTaskAll(userID uint) ([]*GetTasks, error)
	UpdateTask(id uint, userID uint, updateTask Task) error
	DeleteTask(id uint, userID uint) error
}