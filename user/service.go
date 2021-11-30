package user

type Service interface {
	FindByNameProduct(name_product string) ([]User, error)
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