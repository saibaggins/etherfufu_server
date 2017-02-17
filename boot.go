package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/controllers"
	"github.com/saibaggins/etherfufu-server/core"
	"net/http"
)

func main() {
	if core.ActiveENV() != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Define the router
	router := gin.Default()

	// Load Middlewares
	router.Use(core.CORSMiddleware())

	// Define the routes
	v1 := router.Group("/audiobank/v1")
	{
		audiobank_controller := new(controllers.DisplayOptionController)
		metadata_controller := new(controllers.MetadataController)

		v1.GET("options", audiobank_controller.ListAll)
		v1.PUT("options", audiobank_controller.Create)
		v1.POST("metadata", metadata_controller.Create)

	}

	// Handle No routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Path is unavailable",
		})
	})

	// Initialize the server
	http.ListenAndServe(":3000", router)
}
