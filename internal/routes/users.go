package routes

import (
	"Golang-jwt/internal/controllers"
	"Golang-jwt/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterUsers(router *gin.Engine) {
	//authController

	// authrepo := &repositories.UserRepository{}
	authservice := services.UserService{}
	authController := controllers.NewUserController(authservice)

	router.POST("/api/auth/login", authController.LoginAction)
	router.POST("/api/auth/register", authController.RegisterAction)
	router.POST("/api/auth/register", authController.RegisterAction)
}
