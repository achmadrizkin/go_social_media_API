package userfollower

type Service interface {
	JoinUserToUserFollowers(name_user string) ([]UserFollower, error)
	Create(userFollower UserFollowerRequest) (UserFollower, error)
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

func (s *service) Create(userFollowerRequest UserFollowerRequest) (UserFollower, error) {
	book := UserFollower {
		User_id:    userFollowerRequest.User_id,
		Name_user:  userFollowerRequest.Name_user,
		Email_user: userFollowerRequest.Email_user,
		Image_url:  userFollowerRequest.Image_url,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
	// return s.repository.FindAll()
}