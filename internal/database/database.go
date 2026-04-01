package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect () (*sql.DB, error) {

	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatal("DATABASE_URL não encontrada")
	}

	db, err :=  sql.Open("pgx", dbURL)
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