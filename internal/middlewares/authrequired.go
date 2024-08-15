package middlewares

import (
	"fmt"

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
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.JSON(401, map[string]string{
				"error: ": "Missing authorization header",
			})
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := verifyToken(tokenString)

		if err != nil {
			ctx.JSON(401, map[string]string{
				"error: ": "unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
