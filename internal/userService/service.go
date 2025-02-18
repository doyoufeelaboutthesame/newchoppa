package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}
func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}
func (s *UserService) UpdateUserByID(id int64, user User) (User, error) {
	return s.repo.UpdateUserByID(uint(id), user)
}
func (s *UserService) DeleteUserByID(id int64) error {
	return s.repo.DeleteUserByID(uint(id))
}
