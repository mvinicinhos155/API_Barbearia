package database

import (
	"database/sql"
	"fmt"
)

func CreateTableUser (db *sql.DB) error {

	query := `
		CREATE TABLE If NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT'USER',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

		_, err := db.Exec(query)
			if err != nil {
				fmt.Println("Erro ao criar tabela")
				return err
			}
		
		fmt.Println("Tabela criada com sucesso")
		return nil
}

func CreateTableCortes (db *sql.DB) error {

	query := (`CREATE TABLE IF NOT EXISTS haircut (
				id SERIAL PRIMARY KEY,
				haircut_style TEXT NOT NULL,
				price DECIMAL(10, 2),
				day TEXT NOT NULL,
				hour TEXT NOT NULL,
				user_id INT,
				FOREIGN KEY (user_id) REFERENCES users(id)
				);
			`)

		_, err := db.Exec(query)
			if err != nil {
				fmt.Println("Erro ao criar tabela")
				return err
			}

		fmt.Println("Tabela criada com sucesso!!")
		return nil
}

func Migrations(db *sql.DB) {

}