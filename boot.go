package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/controllers"
	"github.com/saibaggins/etherfufu-server/utils"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}

func main() {
	utils.StartDBSession("mongodb://localhost:27017/test")
	session := utils.GetDB()
	defer session.Close()

	fmt.Println(session.DB("test").C("sessions").Count())

	r := gin.Default()

	r.Use(CORSMiddleware())

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		v1.GET("/ping", new(controllers.HomeController).Ping)
		v1.GET("/audiobank/options", new(controllers.AudioBankController).ListAll)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Path is unavailable",
		})
	})

	r.Run(":3000")
}
