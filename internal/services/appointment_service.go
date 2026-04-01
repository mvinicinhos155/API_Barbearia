package services

import (
	"api_barbearia/internal/models"
	"database/sql"
	"fmt"
	"time"
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

func GetAppointByUserId(db *sql.DB, user_id int) ([]models.Appointments, error) {
	query := "SELECT id, user_id, haircut_id, date FROM appointments WHERE user_id = $1"

	rows, err := db.Query(query, user_id)

		if err != nil {
			fmt.Println("Erro com banco de dados", err)
			return nil, err
		}

	defer rows.Close()

	var appoints []models.Appointments

	for rows.Next() {
		var appoint models.Appointments

		err := rows.Scan(&appoint.ID, &appoint.UserId, &appoint.HaircutId, &appoint.Date)
			if err != nil {
				return nil, err
			}

		appoints = append(appoints, appoint)
	}

	return appoints, nil
}

func GetAppointByDate(db *sql.DB, date string) ([]string, error) {
	query := `SELECT date FROM appointments WHERE DATE(date) = $1`

		rows, err := db.Query(query, date)
			if err != nil {
				return nil, err
			}
	defer rows.Close()

	var times []string
	
	for rows.Next(){
		var fulldate time.Time
		err := rows.Scan(&fulldate)
			if err != nil {
				return nil, err
			}

		hour := fulldate.Format("15:04")
		times = append(times, hour) 
	}

	return times, nil
}

func GetAllAppointment (db *sql.DB) ([]models.Appointments, error) {
	query := "SELECT id, user_id, haircut_id, date, notified FROM appointments"

	rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		
	defer rows.Close()

	var appoints []models.Appointments

	for rows.Next() {
		var appoint models.Appointments

		err := rows.Scan(&appoint.ID, &appoint.UserId, &appoint.HaircutId, &appoint.Date, &appoint.Notified)
			if err != nil {
				return nil, err
			}
		
		appoints = append(appoints, appoint)
	}

	return appoints, nil
}