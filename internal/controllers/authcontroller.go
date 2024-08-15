package controllers

import (
	"Golang-jwt/internal/dtos"
	"Golang-jwt/internal/services"
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (l *AuthController) LoginAction(ctx *gin.Context) {

	var dto dtos.LoginDTO
	json.NewDecoder(ctx.Request.Body).Decode(&dto)

	auth, err := l.service.Login(dto.Login, dto.Password)

	if err != nil {
		ctx.JSON(400, map[string]string{
			"mesage": "invalid credentials",
		})
	}

	ctx.JSON(200, auth)
}

func (l *AuthController) RegisterAction(ctx context.Context, login, password string) {

}
