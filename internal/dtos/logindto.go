package dtos

type LoginDTO struct {
	Login    string
	Password string
}

type RegisterDTO struct {
	Email    string
	Fullname string
	Address  string
}
