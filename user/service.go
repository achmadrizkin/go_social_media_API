package user

type Service interface {
	FindByNameProduct(name_product string) ([]User, error)
	FindByEmail(email string) ([]User, error)
	CreateIfNotExistOrUpdateIfExist(email string) ([]User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByNameProduct(name_product string) ([]User, error) {
	users, err := s.repository.FindByNameProduct(name_product)
	return users, err
	// return s.repository.FindAll()
}

func (s *service) FindByEmail(email string) ([]User, error) {
	users, err := s.repository.FindByEmail(email)
	return users, err
	// return s.repository.FindAll()
}

func (s *service) CreateIfNotExistOrUpdateIfExist(email string) ([]User, error) {
	users, err := s.repository.CreateIfNotExistOrUpdateIfExist(email)
	return users, err
	// return s.repository.FindAll()
}
