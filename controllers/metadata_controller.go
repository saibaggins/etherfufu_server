package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/core"
	"github.com/saibaggins/etherfufu-server/models"
	"github.com/saibaggins/etherfufu-server/utils"
	"github.com/satori/go.uuid"
	"net/http"
)

var AudioSampleBucketName = "com.saibaggins.etherfufu." + core.ActiveENV()

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
	s3Session := utils.S3Service()
	metadata_id := c.Param("metadata_id")
	file, _, err := c.Request.FormFile("upload")
	if err == nil {
		_, err := s3Session.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(AudioSampleBucketName),
			Key:         aws.String("/audio_samples/" + metadata_id),
			Body:        file,
			ContentType: aws.String("audio/wav"),
		})

		if err == nil {
			c.AbortWithStatus(http.StatusOK)

		} else {
			fmt.Print(err)
			c.AbortWithStatus(500)
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "CORRUPT_FILE_UPLOAD",
			"message": err.Error(),
		})
	}
}
