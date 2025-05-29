package user

type UserService struct {
	repository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repository: repo,
	}
}

func (us *UserService) GetUsers() ([]User, error) {
	return us.repository.GetUsers()
}

func (us *UserService) CreateUser(user User) (User, error) {
	userID, err := us.repository.CreatUser(user)

	if err != nil {
		return User{}, err
	}

	user.ID = userID

	return user, nil
}

func (us *UserService) GetUserByID(userID int) (*User, error) {
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
