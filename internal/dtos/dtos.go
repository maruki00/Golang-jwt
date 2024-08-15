package dtos

type LoginDTO struct {
	Login    string
	Password string
}

type RegisterDTO struct {
	Email    string
	Fullname string
	Address  string
	Password string
}

type AuthDTO struct {
	Id       int
	Email    string
	Fullname string
	Token    string
}

type GetUsersDTO struct {
	Page   int
	Offset int
}
