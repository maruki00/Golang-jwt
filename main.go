package main

import (
	"Golang-jwt/internal/middlewares"
	"Golang-jwt/internal/routes"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string
	Password string
}

func main() {

	router := gin.New()
	router.Use(middlewares.AuthRequired())
	routes.RegisterUsers(router)

	router.Run(":3000")
}
