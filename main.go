package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Body)
	})
}
