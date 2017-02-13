package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

// Default home controller
type HomeController struct {
}

func (ctrl *HomeController) Ping(c *gin.Context) {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	type Response struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}
	sample := Response{u1.String(), "pong"}
	fmt.Println(sample)
	c.JSON(http.StatusOK, sample)
}
