package explore

import "time"

type Explore struct {
	Id               int
	Name_user        string
	Email_user       string
	Image_url        string

	Image_post       string
	Description_post string
	Like_post        int

	Following        int
	Followers        int
	
	CreatedAt        time.Time
	UpdateAt         time.Time
}
