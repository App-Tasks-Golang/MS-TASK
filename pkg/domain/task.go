package domain

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `gorm:"type:varchar(255);index" json:"title"`
	Description string `json:"description"`
}

type GetTasks struct {
	IdTask      uint   `json:"id_task"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	_           struct{} `gorm:"-"`
}
