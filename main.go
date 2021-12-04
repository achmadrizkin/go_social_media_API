package main

import (
	"log"

	"github.com/achmadrizkin/go_social_media_API/explore"
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

	db.AutoMigrate(&user.User{}, &explore.Explore{})

	// API Versioning
	v1 := r.Group("/v1")

	// USER
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewBookHandler(userService)

	v1.GET("/users/:name_user", userHandler.GetUserByName)

	// EXPLORE
	exploreRepository := explore.NewRepository(db)
	exploreService := explore.NewService(exploreRepository)
	exploreHandler := handler.NewExploreHandler(exploreService)

	v1.POST("/post", exploreHandler.PostExploreHandler)
	v1.GET("/explore/:user", exploreHandler.GetExploreNotUserAndOrderByLike)
	v1.GET("/explore/user/:email", exploreHandler.GetExploreByEmailAndOrderByCreateAt)

	// GET BY EMAIL AND ORDER BY CreateAt http://localhost:3000/v1/explore/user/:email

	r.Run(":3000")
}