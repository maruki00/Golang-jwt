package internal

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func (l *AuthController) LoginAction(ctx *gin.Context) {

}

func (l *AuthController) RegisterAction(ctx context.Context, login, password string) {

}
