package services

import (
	"api_barbearia/internal/models"
	"database/sql"
	"fmt"
)

func InsertMessage (db *sql.DB, message *models.Message) error {

	query := `INSERT INTO messagens(name, email, message) VALUES ($1, $2, $3)`

	_, err := db.Exec(query, &message.Name, &message.Email, &message.Message)
		if err != nil {
			fmt.Println("Erro com o banco de dados")
			return err
		}

	fmt.Println("Valores adcionado com sucesso")
	return nil
}

func GetMessagens(db *sql.DB) ( []models.Message, error) {

	query := `SELECT id, name, email, message FROM messagens`
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println("Erro com banco de dados")
			return  nil, err
		}
	defer rows.Close()

	var messagens []models.Message
	for rows.Next() {
		var message models.Message

		err := rows.Scan(&message.Id, &message.Name, &message.Email, &message.Message)
			if err != nil  {
				return nil, err
			}
		
		messagens = append(messagens, message)
	}

	return messagens, nil
}