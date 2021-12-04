package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindByNameProduct(name_product string) ([]User, error)
	FindByEmail(name_product string) ([]User, error)
	CreateIfNotExistOrUpdateIfExist(email string) ([]User, error)
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

func (r *repository) FindByEmail(email string) ([]User, error) {
	var user []User
	// err := r.db.Where("title = ?", title).First(&users).Error

	err := r.db.Where("email_user LIKE ?", email).Find(&user).Error

	// err := r.db.Where("email_user LIKE ?", name_product).Find(&allProducts).Error

	return user, err
}

func (r *repository) CreateIfNotExistOrUpdateIfExist(email string) ([]User, error) {
	var user []User 
	err := r.db.Where(User{Email_user: email}).Assign(User{Email_user: email}).FirstOrCreate(&User{}).Error

	return user, err
}