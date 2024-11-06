package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/nomannaq/e-commerce-restfulAPI-go/cmd/api"
	"github.com/nomannaq/e-commerce-restfulAPI-go/cmd/config"
	"github.com/nomannaq/e-commerce-restfulAPI-go/cmd/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		DBName:               config.Envs.DBName,
		Addr:                 config.Envs.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	server := api.NewAPIServer(":8080", db)
	log.Println("Listening on port 8080")

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MySQL!")

}
