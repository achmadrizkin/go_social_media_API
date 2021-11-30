package handler

import (
	"net/http"

	"github.com/achmadrizkin/go_social_media_API/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewBookHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUserByName(c *gin.Context) {
	name_user := c.Param("name_user")

	b, err := h.userService.FindByNameProduct(name_user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []user.UserResponse
	for _, b := range b {
		bookResponse := converToAllUserResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	if booksResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": booksResponse,
		})
	}
}

func converToAllUserResponse(b user.User) user.UserResponse {
	return user.UserResponse{
		Id:         b.Id,
		Name_user:  b.Name_user,
		Email_user: b.Email_user,
		Image_url:  b.Image_url,
		Following:  b.Following,
		Followers:  b.Followers,
	}
}
