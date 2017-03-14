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

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

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
	v1 := router.Group("/v1")
	{
		audiobank := v1.Group("/audiobank")
		{
			audiobank_controller := new(controllers.DisplayOptionController)
			audiobank.GET("options", audiobank_controller.ListAll)
			audiobank.PUT("options", audiobank_controller.Create)

			metadata_controller := new(controllers.MetadataController)
			audiobank.POST("metadata", metadata_controller.Create)
			audiobank.POST("metadata/:metadata_id/sample", metadata_controller.UploadAudioSample)
		}

		home_controller := new(controllers.HomeController)
		v1.Any("call", home_controller.Call)

	}

	router.NoRoute(func(c *gin.Context) {
		statusCode := http.StatusNotFound
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    "NOT_FOUND",
			Message: http.StatusText(statusCode),
		})
	})
}
