package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

// Default home controller
type HomeController struct{}

func (ctrl *HomeController) Ping(c *gin.Context) {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	response := struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}{u1.String(), "pong"}

	fmt.Println(response)
	c.JSON(http.StatusOK, response)
}

func (ctrl *HomeController) Call(c *gin.Context) {
	response := struct {
		Color         string  `json:"color"`
		Message       string  `json:"message"`
		Notify        bool    `json:"notify"`
		MessageFormat string  `json:"message_format"`
	}{
		"green",
		"https://meet.jit.si/ether-test",
		true,
		"text",
	}

	c.JSON(http.StatusOK, response)

}
