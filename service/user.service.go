package service

import (
	"github.com/rillmind/apiGin/model"
	"github.com/rillmind/apiGin/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repository: repo,
	}
}

func (us *UserService) GetUsers() ([]model.User, error) {
	return us.repository.GetUsers()
}

// func (us *UserService) x(user model.User) (model.Product, error) {}
// func (us *UserService) x(user model.User) (model.Product, error) {}
// func (us *UserService) x(user model.User) (model.Product, error) {}
