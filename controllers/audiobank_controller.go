package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/models"
	"github.com/satori/go.uuid"
	"net/http"
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
		metadata.ID = uuid.NewV1().String()
		success := new(models.MetadataModel).Put(metadata)
		if success {
			c.JSON(http.StatusOK, metadata)
		} else {
			c.AbortWithStatus(http.StatusConflict)
		}
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
