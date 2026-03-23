package jobs

import (
	"database/sql"
	"fmt"
	"time"
)

func StartReminderJob(db *sql.DB) {
	go func ()  {
		for {
			checkeAppointments(db)
			time.Sleep(1 * time.Minute) 	
		} 
	}()
}

func checkeAppointments(db *sql.DB) {
	query := "SELECT id, user_id, haircut_id, date FROM appointments"

	rows, err := db.Query(query)
		if err != nil {
			fmt.Println("Erro ao bascar agendamento", err)
			return
		}

	defer rows.Close()

	now := time.Now()

	for rows.Next() {
		var id, userID, haircutID int
		var date time.Time

		err := rows.Scan(&id, &userID, &haircutID, &date)
			if err != nil {
				continue
			}

	diff := date.Sub(now)

	for diff > 0 && diff <= 1*time.Minute {
		fmt.Printf("lembrete: Usuário %d tem um corte em breve", userID)
	}

	}
}