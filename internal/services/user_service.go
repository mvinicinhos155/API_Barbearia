package services

import (
	"database/sql"
	"fmt"
	"api_barbearia/internal/models"
)

func InsertUser(db *sql.DB, user *models.Users) error {

	query := `INSERT INTO users (name, email, password, phone) 
			  VALUES ($1, $2, $3, $4);`

	_, err := db.Exec(query, user.Name, user.Email, user.Password, user.Phone)
	if err != nil {
		fmt.Println("Erro com o banco de dados")
		return err
	}

	return nil
}

func GetUserbyEmail (db *sql.DB, email string) (models.Users , error) {

	var user models.Users
	query := "SELECT id, name, email, password, role FROM users WHERE email = $1"

	 err := db.QueryRow(query, email).Scan(&user.ID, &user.Name,&user.Email, &user.Password, &user.Role)
		if err != nil {
			fmt.Println("Erro com banco de dados", err)
			return user, err
		}

		return user, nil
}

func GetAllUser (db *sql.DB) ([]models.Users, error) {
	query := "SELECT id, name, email, role, phone FROM users"

	rows, err := db.Query(query)
		if err != nil {
			fmt.Println("Erro com banco de dados")
			return nil, err
		}
	
	defer rows.Close()

	var users []models.Users

	for rows.Next() {
		var user  models.Users

		err := rows.Scan(&user.ID, &user.Name, &user.Email,&user.Role, &user.Phone)
			if err != nil {
				return nil, err
			}

		users = append(users, user)
	}

	return users, nil
}

func DeleteUser (db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := db.Exec(query, id)
		if err != nil {
			fmt.Println("Erro com banco de dados", err)
			return err
		}

	fmt.Println("ok..")
	return nil
}