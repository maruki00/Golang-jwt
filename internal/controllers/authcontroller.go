package controllers

import (
	"Golang-jwt/internal/dtos"
	"Golang-jwt/internal/services"
	"encoding/json"
	"fmt"

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
			"mesage": "invalid credentials " + err.Error(),
		})
		return
	}

	ctx.JSON(200, auth)
}

func (l *AuthController) RegisterAction(ctx *gin.Context) {

	var dto dtos.RegisterDTO
	json.NewDecoder(ctx.Request.Body).Decode(&dto)

	fmt.Println("level : ", dto)

	res, err := l.service.Register(dto.Email, dto.Password, dto.Fullname, dto.Address)

	if err != nil {
		ctx.JSON(400, map[string]string{
			"mesage": "invalid credentials",
		})
		return
	}

	ctx.JSON(200, res)
}
