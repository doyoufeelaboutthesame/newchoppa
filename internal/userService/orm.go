package userService

import (
	"TheRealOne/internal/taskService"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []taskService.Task
}
