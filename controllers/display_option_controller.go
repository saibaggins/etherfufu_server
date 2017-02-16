package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/models"
	"net/http"
)

// Default home controller
type DisplayOptionController struct {
}

func (ctrl *DisplayOptionController) ListAll(c *gin.Context) {
	model := new(models.DisplayOptionModel)
	c.JSON(http.StatusOK, model.FindAll())
}

func (ctrl *DisplayOptionController) Create(c *gin.Context) {
	display_option_model := new(models.DisplayOptionModel)

	var display_options []models.DisplayOption
	if c.BindJSON(&display_options) == nil && display_option_model.Put(display_options) {
		c.JSON(http.StatusOK, display_options)
	} else {
		c.AbortWithStatus(http.StatusConflict)
	}

}
