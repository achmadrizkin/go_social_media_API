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