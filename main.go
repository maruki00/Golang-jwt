package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string
	Password string
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {

		fmt.Println(ctx.Request.Body)
		var u User
		json.NewDecoder(ctx.Request.Body).Decode(&u)
		ctx.JSON(200, u)
	})
	router.Run()
}
