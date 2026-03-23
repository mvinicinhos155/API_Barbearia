package services

import (
	"database/sql"
	"fmt"
	"api_barbearia/internal/models"
)

func InsertHaircut (db *sql.DB, hair *models.Haircuts) error {
	query := "INSERT INTO haircuts (name, price) VALUES ($1, $2)"

	_, err := db.Exec(query, hair.Name, hair.Price)
		if err != nil {
			fmt.Println("Erro ao enviar o tipo de corte")
			return err
		}

	fmt.Println("Corte enviado com sucesso")
	return nil
}

func GetAllHairs(db *sql.DB) ([]models.Haircuts ,error) {

	query := "SELECT * FROM haircuts"

	rows, err := db.Query(query)
		if err != nil {
			fmt.Println("Erro com o banco de dados")
			return nil, err
		}

	defer rows.Close()

		var hairs []models.Haircuts

	for rows.Next() {
		var hair models.Haircuts

		err := rows.Scan(&hair.ID, &hair.Name, &hair.Price)
			if err != nil {
				return nil, err
			}

		hairs = append(hairs, hair)
	}

	return hairs, nil
}

func DeleteHair(db *sql.DB, id int) error {
	query := "DELETE FROM haircuts WHERE id = $1"

	_, err := db.Exec(query, id)
		if err != nil {
			fmt.Println("Erro com banco de dados", err)
			return err
		}

	fmt.Println("ok..")
	return nil
}

