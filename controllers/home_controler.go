package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (ctrl *HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"message":    "Welcome to the Home Page",
		"categories": [2]string{"Hello", "World"},
	})
}
