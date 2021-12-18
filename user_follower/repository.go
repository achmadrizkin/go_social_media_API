package userfollower

import (
	"gorm.io/gorm"
)

type Repository interface {
	JoinUserToUserFollowers(name_user string) ([]UserFollower, error)
	Create(userFollower UserFollower) (UserFollower, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) JoinUserToUserFollowers(name_user string) ([]UserFollower, error) {
	var comment []UserFollower

	// err := r.db.Where("").Find(&comment).Error
	// err := r.db.Model(&UserFollower{}).Select("name_user, email_user").Joins("left join users on name_user = users.user_id").Where("name_user = ?", name_user).Scan(&User{}).Error

	//
	err := r.db.Table("user_followers").Select("user_followers.name_user, user_followers.email_user, user_followers.image_url, user_followers.user_id, user_followers.id").Joins("left join users on users.email_user = user_followers.user_id").Where("users.email_user = ?", name_user).Find(&comment).Error

	return comment, err
}

func (r *repository) Create(userFollower UserFollower) (UserFollower, error) {
	err := r.db.Create(&userFollower).Error

	return userFollower, err
}
