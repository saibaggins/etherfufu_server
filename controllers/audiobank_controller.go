package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/models"
	"net/http"
	"github.com/satori/go.uuid"
)

// Default home controller
type AudioBankController struct {
}

func (ctrl *AudioBankController) ListAll(c *gin.Context) {
	displayOptionModel := models.DisplayOptionModel{}
	c.JSON(http.StatusOK, displayOptionModel.FindAll())
}

func (self *AudioBankController) CreateMetadata(c *gin.Context) {
	var metadata models.Metadata
	if c.BindJSON(&metadata) == nil {
		metadata.ID = uuid.NewV1()
		c.JSON(http.StatusOK, metadata)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status" : "bad request"})
	}
}
