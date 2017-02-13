package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/models"
	"net/http"
)

// Default home controller
type AudioBankController struct {
}

func (ctrl *AudioBankController) ListAll(c *gin.Context) {
	selectionOptionModel := models.SelectionOptionModel{}
	c.JSON(http.StatusOK, selectionOptionModel.FindAll())
}
