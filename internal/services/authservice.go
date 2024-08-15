package services

import (
	"Golang-jwt/internal/dtos"
	"Golang-jwt/internal/repositories"
)

type AuthService struct {
	repo repositories.UserRepository
}

func (l *AuthService) Login(login, password string) (*dtos.AuthDTO, error) {

	return l.repo.Login(login, password)
}

func (l *AuthService) Register(email, password, fullname, address string) (*dtos.RegisterDTO, error) {

	return l.repo.Register(email, fullname, address, password)
}
