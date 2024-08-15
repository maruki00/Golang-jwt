package main

import (
	"Golang-jwt/internal/routes"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string
	Password string
}

func main() {

	router := gin.Default()
	routes.RegisterAuth(router)

	router.Run(":3000")
}
