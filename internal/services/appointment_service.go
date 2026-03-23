package services

import (
	"database/sql"
	"fmt"
	"api_barbearia/internal/models"
)

func InsertAppointment(db *sql.DB, appoint *models.Appointments) error {
	query := "INSERT INTO appointments (user_id, haircut_id, date) VALUES ($1, $2, $3)"

	_, err := db.Exec(query, appoint.UserId, appoint.HaircutId, appoint.Date)
		if err != nil {
			fmt.Println("Erro com banco de dados", err)
			return err
		}

	return nil
}

func GetAppointByUserId(db *sql.DB, user_id int) (*models.Appointments, error) {
	query := "SELECT id, user_id, haircut_id, date FROM appointments WHERE user_id = $1"

	var apoint models.Appointments

	err := db.QueryRow(query, user_id).Scan(&apoint.ID, &apoint.UserId, &apoint.HaircutId, &apoint.Date)
		if err != nil {
			fmt.Println("Erro com banco de dados", err)
			return &apoint, err
		}

	return &apoint, nil
}