package main

import (
	"uooobarry/gin-notes/config"
	"uooobarry/gin-notes/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./public")

	routes.Setup(r)

	r.Run(":8080")
}
