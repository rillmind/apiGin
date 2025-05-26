package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	docker = "go_db"
	local  = "127.0.0.1"

	host     = local
	port     = 5432
	password = "1234"
	user     = "postgres"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)
	return db, nil
}
