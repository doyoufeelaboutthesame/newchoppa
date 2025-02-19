package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (repo *taskRepository) CreateTask(task Task) (Task, error) {
	result := repo.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}
func (repo *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := repo.db.Find(&tasks).Error
	return tasks, err
}
func (repo *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var oldTask Task
	res := repo.db.First(&oldTask, id)
	if res.Error != nil {
		return Task{}, res.Error
	}
	oldTask.Task = task.Task
	oldTask.IsDone = task.IsDone
	oldTask.UserID = task.UserID
	res = repo.db.Save(&oldTask)
	if res.Error != nil {
		return Task{}, res.Error
	}
	return oldTask, nil
}
func (repo *taskRepository) DeleteTaskByID(id uint) error {
	res := repo.db.Delete(&Task{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
