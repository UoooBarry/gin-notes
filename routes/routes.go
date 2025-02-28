package routes

import (
	"uooobarry/gin-notes/config"
	"uooobarry/gin-notes/controllers"
	"uooobarry/gin-notes/repositories"
	"uooobarry/gin-notes/services"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	router.POST("/users", userController.CreateUser)
	// router.POST("/login", controllers.LoginUser)

	homeController := controllers.NewHomeController()
	router.GET("/", homeController.Index)
}
