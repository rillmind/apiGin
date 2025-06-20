package auth

import (
	"errors"

	"github.com/rillmind/apiGin/hash"
	"github.com/rillmind/apiGin/jwt"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		repository,
	}
}

func (as *Service) Login(login, password string) (string, error) {
	jwtService := jwt.NewService()

	user, err := as.repository.Login(login)

	if err != nil {
		return "", nil
	}

	if user == nil {
		return "", errors.New("usuário não encontrado")
	}

	if !hash.CheckPassword(password, user.Password) {
		return "", errors.New("senha inválida")
	}

	token, err := jwtService.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
