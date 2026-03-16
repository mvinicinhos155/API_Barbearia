package services

import (
	"database/sql"
	"fmt"
	"api_barbearia/internal/models"
)

func InsertHairCut (db *sql.DB, hair *models.Haircuts) error {


	query := "INSERT INTO haircut (haircut_style, price, day, hour, user_id) VALUES ($1, $2, $3, $4, $5);"

     _ ,err := db.Exec(query, hair.Haircut_style, hair.Price, hair.Day, hair.Hour, hair.User_id)
	  if err != nil {
		fmt.Println("Erro com banco de dados")
		return err
	  }

	fmt.Println("dados criado com sucesso!")
	return nil

}