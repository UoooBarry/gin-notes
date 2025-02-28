package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderErrorPage(status int, c *gin.Context, err error) {
	var temp string
	message := err.Error()
	switch status {
	case http.StatusBadRequest:
		temp = "400.html"
	default:
		temp = "500.html"
		message = "Internal Server Error"
	}
	c.HTML(status, temp, gin.H{
		"message": message,
	})
}
