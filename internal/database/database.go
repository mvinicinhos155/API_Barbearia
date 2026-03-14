package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect () (*sql.DB, error) {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)


	db, err :=  sql.Open("pgx", dsn)
		if err != nil {
			return nil, err
		}

	err = db.Ping()
		if err != nil {
			return nil, err
		}

	log.Println("Banco de dados conectado")

	return db, nil
}