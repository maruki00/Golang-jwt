package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("golang-jwt")

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	fmt.Println("tokn : ", token)
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func AuthRequired() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" || len(tokenString) < 8 {
			ctx.AbortWithStatusJSON(401, map[string]string{
				"error: ": "Missing authorization header",
			})
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 7)

		err := verifyToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(401, map[string]string{
				"error: ": "unauthorized",
			})

		}
		ctx.Next()
	}
}
