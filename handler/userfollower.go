package handler

import (
	"fmt"
	"net/http"

	"github.com/achmadrizkin/go_social_media_API/user_follower"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userFollowerHandler struct {
	userfollowerService userfollower.Service
}

func NewUserFollowerHandler(userfollowerService userfollower.Service) *userFollowerHandler {
	return &userFollowerHandler{userfollowerService}
}

func (h *userFollowerHandler) PostUserFollowerHandler(c *gin.Context) {
	var userFollowerRequest userfollower.UserFollowerRequest

	err := c.ShouldBindJSON(&userFollowerRequest)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	hoodie, err := h.userfollowerService.Create(userFollowerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hoodie,
	})
}


func (h *userFollowerHandler) GetJoinUserToUserFollowers(c *gin.Context) {
	user := c.Param("user")

	allproductss, err := h.userfollowerService.JoinUserToUserFollowers(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []userfollower.UserFollowerResponse

	for _, b := range allproductss {
		allproductsResponse := converToUserFollowersResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}

func converToUserFollowersResponse(b userfollower.UserFollower) userfollower.UserFollowerResponse {
	return userfollower.UserFollowerResponse{
		Id:         b.Id,
		User_id:    b.User_id,
		Name_user:  b.Name_user,
		Email_user: b.Email_user,
		Image_url:  b.Image_url,

		CreatedAt:  b.CreatedAt,
		UpdateAt:   b.UpdateAt,
	}
}
