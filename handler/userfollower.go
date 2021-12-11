package handler

import (
	"net/http"

	"github.com/achmadrizkin/go_social_media_API/user_follower"
	"github.com/gin-gonic/gin"
)

type userFollowerHandler struct {
	userfollowerService userfollower.Service
}

func NewUserFollowerHandler(userfollowerService userfollower.Service) *userFollowerHandler {
	return &userFollowerHandler{userfollowerService}
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
