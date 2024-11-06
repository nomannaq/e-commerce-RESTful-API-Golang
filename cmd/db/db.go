package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/mydatabase")

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
