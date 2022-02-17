package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/isyarah/models"
	"github.com/madjiebimaa/isyarah/requests"
)

type LocationHandler struct {
	locationService models.LocationService
}

func NewLocationHandler(locationService models.LocationService) *LocationHandler {
	return &LocationHandler{
		locationService,
	}
}

func (l *LocationHandler) Create(c *gin.Context) {
	var req requests.LocationCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body request",
		})
		return
	}

	ctx := c.Request.Context()
	location, err := l.locationService.Create(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "there's error in service",
		})
		return
	}

	c.JSON(http.StatusOK, location)
}
