package handler

import (
	"fmt"
	"net/http"

	"github.com/achmadrizkin/go_social_media_API/explore"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type exploreHandler struct {
	exploreService explore.Service
}

func NewExploreHandler(hoodieService explore.Service) *exploreHandler {
	return &exploreHandler{hoodieService}
}

func (h *exploreHandler) PostExploreHandler(c *gin.Context) {
	var hoodieRequest explore.ExploreRequest

	err := c.ShouldBindJSON(&hoodieRequest)

	if err != nil {
		// log.Fatal(err) -> kalau terjadi error, server mati
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)

			// gunakan return untuk tidak melanjutkan yang dibawah
			return
		}
	}

	hoodie, err := h.exploreService.Create(hoodieRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": hoodie,
	})
}

func (h *exploreHandler) GetExploreNotUserAndOrderByLike(c *gin.Context) {
	user := c.Param("user")

	allproductss, err := h.exploreService.FindByNotUserAndOrderByLike(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []explore.ExploreResponse

	for _, b := range allproductss {
		allproductsResponse := converToExploreResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}

func (h *exploreHandler) GetExploreByEmailAndOrderByCreateAt(c *gin.Context) {
	email := c.Param("email")

	allproductss, err := h.exploreService.GetByEmailAndOrderByCreateAt(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []explore.ExploreResponse

	for _, b := range allproductss {
		allproductsResponse := converToExploreResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}

func (h *exploreHandler) GetUserFollowingPost(c *gin.Context) {
	email := c.Param("email")

	allproductss, err := h.exploreService.GetByUserFollowing(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []explore.ExploreResponse

	for _, b := range allproductss {
		allproductsResponse := converToExploreResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}

func converToExploreResponse(b explore.Explore) explore.ExploreResponse {
	return explore.ExploreResponse{
		Id:               b.Id,
		Name_user:        b.Name_user,
		Email_user:       b.Email_user,
		Image_url:        b.Image_url,
		Image_post:       b.Image_post,
		Description_post: b.Description_post,
		Like_post:        b.Like_post,
		CreatedAt:        b.CreatedAt,
		UpdateAt:         b.UpdateAt,
	}
}
