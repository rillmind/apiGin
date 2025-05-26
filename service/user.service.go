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

func (us *UserService) CreateUser(user model.User) (model.User, error) {
	userID, err := us.repository.CreatUser(user)

	if err != nil {
		return model.User{}, err
	}

	user.ID = userID

	return user, nil
}

func (us *UserService) GetUserByID(userID int) (*model.User, error) {
	user, err := us.repository.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) DeleteUserByID(userID int) (int64, error) {
	user, err := us.repository.DeleteUserByID(userID)

	if err != nil {
		return 0, err
	}

	return user, err
}
