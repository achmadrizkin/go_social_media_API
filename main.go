package main

import (
	"log"

	"github.com/achmadrizkin/go_social_media_API/handler"
	"github.com/achmadrizkin/go_social_media_API/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// connect to database
	// PLEASE CREATE go-ecommerce database first.
	dsn := "root:@tcp(127.0.0.1:3306)/go-social-media?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&user.User{})

	// API Versioning
	v1 := r.Group("/v1")

	//
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewBookHandler(userService)

	v1.GET("/users/:name_user", userHandler.GetUserByName)

	r.Run(":3000")
}