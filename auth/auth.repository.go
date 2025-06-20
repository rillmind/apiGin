package auth

import (
	"database/sql"
	"errors"

	"github.com/rillmind/apiGin/user"
)

type Repository struct {
	connection *sql.DB
}

func NewRepository(connection *sql.DB) Repository {
	return Repository{
		connection,
	}
}

func (ar *Repository) Login(login string) (*user.Model, error) {
	var user user.Model
	query := `
			SELECT
				id,
				user_name,
				user_username,
				user_email,
				user_password,
				user_role
			FROM "user"
			WHERE user_username = $1 OR user_email = $1
		`
	err := ar.connection.QueryRow(query, login).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}

	return &user, nil
}
