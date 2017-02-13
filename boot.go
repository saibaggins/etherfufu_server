package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saibaggins/etherfufu-server/controllers"
	"github.com/saibaggins/etherfufu-server/utils"
	"net/http"
)

func main() {
	utils.StartDBSession("mongodb://localhost:27017/test")
	session := utils.GetDB()
	defer session.Close()

	result, _ := session.DB("test").C("sessions").Count()
	fmt.Println(result)

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
