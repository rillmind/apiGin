package user

type Service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return Service{
		repository: repo,
	}
}

func (us *Service) GetUsers() ([]Model, error) {
	return us.repository.GetUsers()
}

func (us *Service) CreateUser(user Model) (Model, error) {
	userID, err := us.repository.CreatUser(user)

	if err != nil {
		return Model{}, err
	}

	user.ID = userID

	return user, nil
}

func (us *Service) GetUserByID(userID int) (*Model, error) {
	user, err := us.repository.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *Service) DeleteUserByID(userID int) (int64, error) {
	user, err := us.repository.DeleteUserByID(userID)

	if err != nil {
		return 0, err
	}

	return user, err
}
