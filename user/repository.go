package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindByNameProduct(name_product string) ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByNameProduct(name_user string) ([]User, error) {
	var user []User
	// err := r.db.Where("title = ?", title).First(&users).Error

    value := fmt.Sprintf("%%%s%%", name_user)
	err := r.db.Where("name_user LIKE ?", value).Find(&user).Error

	// err := r.db.Where("email_user LIKE ?", name_product).Find(&allProducts).Error

	return user, err
}