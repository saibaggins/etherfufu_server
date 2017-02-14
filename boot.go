package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/controllers"
	"github.com/saibaggins/etherfufu-server/utils"
	"net/http"
)

func main() {
	// Initialize the configuration
	config := utils.LoadEnvConfig()
	fmt.Println(config)

	// Initialize the database
	//////////////////////////////////////////////////////
	// utils.StartDBSession(config.Database.URI)	    //
	// session := utils.GetDB()			    //
	// defer session.Close()			    //
	//////////////////////////////////////////////////////

	// Define the router
	r := gin.Default()

	// Load Middlewares
	r.Use(CORSMiddleware())

	// Define the routes
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", new(controllers.HomeController).Ping)
		v1.GET("/audiobank/options", new(controllers.AudioBankController).ListAll)
	}

	// Handle No routes
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Path is unavailable",
		})
	})

	// Initialize the server
	r.Run(":3000")
}
