package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/controllers"
	"github.com/saibaggins/etherfufu-server/core"
	"github.com/saibaggins/etherfufu-server/core/middlewares"
	"net/http"
)

const (
	Production = "production"
)

func main() {

	setAppVerbosity()

	router := gin.New()
	loadMiddlewares(router)
	defineRoutes(router)

	// Initialize the server
	http.ListenAndServe(":3000", router)
}

func setAppVerbosity() {
	if core.ActiveENV() == Production {
		gin.SetMode(gin.ReleaseMode)
	}
}

// Load Middlewares
func loadMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.ErrorLogger())
}

// Define the routes
func defineRoutes(router *gin.Engine) {
	v1 := router.Group("/audiobank/v1")
	{
		audiobank_controller := new(controllers.DisplayOptionController)
		metadata_controller := new(controllers.MetadataController)

		v1.GET("options", audiobank_controller.ListAll)
		v1.PUT("options", audiobank_controller.Create)
		v1.POST("metadata", metadata_controller.Create)
		v1.POST("metadata/:metadata_id/sample", metadata_controller.UploadAudioSample)

	}

	router.NoRoute(func(c *gin.Context) {
		statusCode := http.StatusNotFound
		c.JSON(http.StatusNotFound, gin.H{
			"code":    statusCode,
			"message": http.StatusText(statusCode),
		})
	})
}
