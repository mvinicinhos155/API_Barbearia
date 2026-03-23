package handler

import (
	"api_barbearia/internal/models"
	"api_barbearia/internal/services"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func HandlerCreateAppointment(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	userID := r.Context().Value("userID").(int)

	var Input struct {
		HaircutId int `json:"haircut_id"`
		Date time.Time `json:"date"`
	}

	err := json.NewDecoder(r.Body).Decode(&Input)
		if err != nil {
			http.Error(w, "Erro ao enviar dados para banco de dados", http.StatusBadRequest)
			return
		}

	appointment := models.Appointments {
		UserId: userID,
		HaircutId: Input.HaircutId,
		Date: Input.Date,
	}

	err = services.InsertAppointment(db, &appointment)
		if err != nil {
			http.Error(w, "Erro com banco de dados", http.StatusInternalServerError)
			return
		}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Agendamento feito com sucesso, espere ser chamado",
		"appointment" : appointment,
	})
}

func HandlerGetByUserId(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userId := r.Context().Value("userID").(int)

	appoint, err := services.GetAppointByUserId(db, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "agendamento do usuário listado com sucesso",
		"appointment" : appoint,
	})
}