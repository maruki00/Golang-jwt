package internal

type AuthService struct {
	repo UserRepository
}

func (l *AuthService) Login(login, password string) (*AuthDTO, error) {

	return l.repo.Login(login, password)
}

func (l *AuthService) Register(login, password string) (*AuthDTO, error) {

	return l.repo.Login(login, password)

}
