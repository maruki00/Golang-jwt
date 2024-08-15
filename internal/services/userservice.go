package services

import (
	"Golang-jwt/internal/dtos"
	"Golang-jwt/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func (l *UserService) Login(login, password string) (*dtos.AuthDTO, error) {

	return l.repo.Login(login, password)
}

func (l *UserService) Register(email, password, fullname, address string) (*dtos.RegisterDTO, error) {

	return l.repo.Register(email, fullname, address, password)
}
