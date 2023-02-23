package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-rest-without-framework/helpers"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "belajar-sql"
)

func DatabaseConnectoon() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	helpers.PanicIfError(err)

	err = db.Ping()
	helpers.PanicIfError(err)

	log.Println("Connected to database!!")

	return db
}
