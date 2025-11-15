package people

type Service interface {
	GetAll() ([]Person, error)
	GetByID(id uint) (Person, error)
	Create(p *Person) error
	Update(p *Person) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]Person, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id uint) (Person, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(p *Person) error {
	return s.repo.Create(p)
}

func (s *service) Update(p *Person) error {
	return s.repo.Update(p)
}

func (s *service) Delete(id uint) error {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(&p)
}
