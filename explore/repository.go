package explore

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(explore Explore) (Explore, error)
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
