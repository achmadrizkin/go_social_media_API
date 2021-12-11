package userfollower

type Service interface {
	JoinUserToUserFollowers(name_user string) ([]UserFollower, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) JoinUserToUserFollowers(name_user string) ([]UserFollower, error) {
	books, err := s.repository.JoinUserToUserFollowers(name_user)
	return books, err
}