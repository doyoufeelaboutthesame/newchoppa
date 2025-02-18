package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user User) (User, error) {
	result := repo.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
func (repo *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := repo.db.Find(&users).Error
	return users, err
}
func (repo *userRepository) DeleteUserByID(id uint) error {
	err := repo.db.Delete(&User{}, id).Error
	return err
}
func (repo *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var oldUser User
	err := repo.db.First(&oldUser, id).Error
	if err != nil {
		return oldUser, err
	}
	oldUser.Email = user.Email
	oldUser.Password = user.Password
	err = repo.db.Save(&oldUser).Error
	if err != nil {
		return oldUser, err
	}
	return oldUser, nil
}
