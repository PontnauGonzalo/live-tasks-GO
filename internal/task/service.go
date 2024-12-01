package task

// connection between cotroller and repository
type Service struct {
	repo *Repository
}

func NewTaskService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (serv *Service) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	return serv.repo.Create(ctx, task)
}