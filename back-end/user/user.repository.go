package user

import (
	"database/sql"
	"fmt"

	"github.com/rillmind/apiGin/hash"
)

type Repository struct {
	connection *sql.DB
}

func NewRepository(connection *sql.DB) Repository {
	return Repository{
		connection,
	}
}

func (ur *Repository) GetUsers() ([]Model, error) {
	var userList []Model
	var userObj Model

	query := `select * from "user"`
	rows, err := ur.connection.Query(query)

	if err != nil {
		fmt.Print(err)
		return []Model{}, err
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
			return []Model{}, err
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil
}

func (ur *Repository) CreatUser(user Model) (int, error) {
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

	hashedPass, err := hash.Password(user.Password)

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

func (ur *Repository) GetUserByID(userID int) (*Model, error) {
	var user Model

	query, err := ur.connection.Prepare(`
		select id, user_name, user_username, user_email, user_password
		from "user"
		where id = $1
	`)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer query.Close()

	err = query.QueryRow(userID).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &user, nil
}

func (ur *Repository) DeleteUserByID(userID int) (int64, error) {
	query, err := ur.connection.Prepare(`
		delete from "user"
		where id = $1
	`)

	if err != nil {
		fmt.Print(err)
		return 0, nil
	}

	defer query.Close()

	result, err := query.Exec(userID)

	if err != nil {
		fmt.Print(err)
		return 0, nil
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Print(err)
		return 0, nil
	}

	return rowsAffected, nil
}
