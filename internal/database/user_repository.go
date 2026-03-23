package database

import (
	"database/sql"
	"fmt"
	"os"
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

func CreateTableAppointment (db *sql.DB) error {
	query := (`CREATE TABLE IF NOT EXISTS appointments(
				id SERIAL PRIMARY KEY,
				user_id INT,
				haircut_id INT,
				date TIMESTAMP NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users(id),
				FOREIGN KEY (haircut_id) REFERENCES haircuts(id) 	
	);`)

	_, err := db.Exec(query)
		if err != nil {
			fmt.Println("Erro com o banco de dados")
			return err
		}

	fmt.Println("Criado com sucesso")
	return nil
}

func UpdateUserRole(db *sql.DB) error {
	var email = os.Getenv("USER_ADMIN")
	fmt.Println(email)
	
	query := `UPDATE users SET role = 'ADMIN' WHERE email = $1 `
	_, err := db.Exec(query, email)
		if err != nil {
			fmt.Println("Erro ao atualizar admin", err)
			return err
		}

	fmt.Println("Atualizado com sucesso")
	return nil
}

func Migrations(db *sql.DB) {
}