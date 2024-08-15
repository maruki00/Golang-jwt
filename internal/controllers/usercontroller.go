package controllers

import (
	"Golang-jwt/internal/dtos"
	"Golang-jwt/internal/services"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (l *UserController) LoginAction(ctx *gin.Context) {

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

func (l *UserController) RegisterAction(ctx *gin.Context) {

	var dto dtos.RegisterDTO
	json.NewDecoder(ctx.Request.Body).Decode(&dto)

	res, err := l.service.Register(dto.Email, dto.Password, dto.Fullname, dto.Address)

	if err != nil {
		ctx.JSON(400, map[string]string{
			"mesage": "invalid credentials",
		})
		return
	}

	ctx.JSON(200, res)
}

func (l *UserController) GetUsersAction(ctx *gin.Context) {

	var dto dtos.GetUsersDTO
	json.NewDecoder(ctx.Request.Body).Decode(&dto)

	res, err := l.service.GetUsers(dto.Page, dto.Offset)

	if err != nil {
		ctx.JSON(400, map[string]string{
			"mesage": "invalid credentials",
		})
		return
	}

	ctx.JSON(200, res)
}
