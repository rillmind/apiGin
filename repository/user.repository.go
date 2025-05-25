package repository

import (
	"database/sql"
	"fmt"

	"github.com/rillmind/apiGin/model"
	"github.com/rillmind/apiGin/utils"
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

func (ur *UserRepository) CreatUser(user model.User) (int, error) {
	var id int

	query, err := ur.connection.Prepare(`
			insert into "user"
			(user_name, user_username, user_email, user_password)
			values ($1, $2, $3, $4) returning id
	`)

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	hashedPass, err := utils.HashPassword(user.Password)

	if err != nil {
		fmt.Print(err)
		return 0, nil
	}

	err = query.QueryRow(user.Name, user.Username, user.Email, hashedPass).Scan(&id)

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (ur *UserRepository) GetUserByID() {}

func (ur *UserRepository) DeleteUserByID() {}
