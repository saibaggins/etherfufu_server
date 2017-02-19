package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/controllers"
	"github.com/saibaggins/etherfufu-server/core"
	"github.com/saibaggins/etherfufu-server/core/middlewares"
	"net/http"
)

const (
	Development = "development"
)

func main() {

	if core.ActiveENV() != Development {
		gin.SetMode(gin.ReleaseMode)
	}

	// Define the router
	router := gin.New()

	// Load Middlewares
	loadMiddlewares(router)

	// Define the routes
	defineRoutes(router)

	// Initialize the server
	http.ListenAndServe(":3000", router)
}

func loadMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.ErrorLogger())
}

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

	// Handle No routes
	router.NoRoute(func(c *gin.Context) {
		statusCode := http.StatusNotFound
		c.JSON(http.StatusNotFound, gin.H{
			"code":    statusCode,
			"message": http.StatusText(statusCode),
		})
	})
}
