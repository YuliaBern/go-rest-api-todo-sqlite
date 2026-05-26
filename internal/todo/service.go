package todo

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(task Task) (Task, error) {
	return s.repo.Create(task)
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repo.GetAll()
}