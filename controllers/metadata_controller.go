package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/models"
	"github.com/satori/go.uuid"
	"net/http"
)

type MetadataController struct{}

func (self *MetadataController) Create(c *gin.Context) {
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

func (self *MetadataController) UploadAudioSample(c *gin.Context) {
	//metadata_id := c.Param("metadata_id")
	//file, header, err := c.Request.FormFile("upload")

}
