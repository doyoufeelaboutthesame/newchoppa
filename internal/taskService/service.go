package taskService

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}
func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}
func (s *TaskService) UpdateTaskByID(id int64, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(uint(id), task)
}
func (s *TaskService) DeleteTaskByID(id int64) error {
	return s.repo.DeleteTaskByID(uint(id))
}
