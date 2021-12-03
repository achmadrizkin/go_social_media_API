package explore

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(explore Explore) (Explore, error)
	FindByNotUserAndOrderByLike(user string) ([]Explore, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(explore Explore) (Explore, error) {
	err := r.db.Create(&explore).Error

	return explore, err
}

func (r *repository) FindByNotUserAndOrderByLike(user string) ([]Explore, error) {
	var explore []Explore

	err := r.db.Where("email_user NOT LIKE ? ORDER BY like_post DESC", user).Find(&explore).Error

	return explore, err
}