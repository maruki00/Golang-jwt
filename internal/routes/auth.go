package routes

import (
	"Golang-jwt/internal/controllers"
	"Golang-jwt/internal/middlewares"
	"Golang-jwt/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterAuth(router *gin.Engine) {
	//authController

	// authrepo := &repositories.UserRepository{}
	authservice := services.AuthService{}
	authController := controllers.NewAuthController(authservice)

	router.POST("/api/auth/login", authController.LoginAction, middlewares.AuthRequired())
}
