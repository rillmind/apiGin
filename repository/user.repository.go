package repository

import (
	"database/sql"
	"fmt"

	"github.com/rillmind/apiGin/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	var userList []model.User
	var userObj model.User

	query := `select * from "user"`
	rows, err := ur.connection.Query(query)

	if err != nil {
		fmt.Print(err)
		return []model.User{}, err
	}

	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Username,
			&userObj.Email,
			&userObj.Password,
		)

		if err != nil {
			fmt.Print(err)
			return []model.User{}, err
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil
}

func (ur *UserRepository) CreatUsers() {}

func (ur *UserRepository) GetUserByID() {}

func (ur *UserRepository) DeleteUserByID() {}
