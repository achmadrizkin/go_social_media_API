package explore

type Service interface {
	Create(exploreRequest ExploreRequest) (Explore, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(exploreRequest ExploreRequest) (Explore, error) {
	following, _ := exploreRequest.Following.Int64()
	followers, _ := exploreRequest.Followers.Int64()

	book := Explore {
		Name_user: exploreRequest.Name_user,
		Email_user: exploreRequest.Email_user,
		Image_url: exploreRequest.Image_url,

		Image_post: exploreRequest.Image_post,
		Description_post: exploreRequest.Description_post,
		Like_post: exploreRequest.Like_post,

		Following: int(following),
		Followers: int(followers),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
	// return s.repository.FindAll()
}