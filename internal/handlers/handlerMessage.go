package handler

import (
	"api_barbearia/internal/models"
	"api_barbearia/internal/services"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandlerSendMessage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var message models.Message

	 err := json.NewDecoder(r.Body).Decode(&message)
	 	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
		}

	err = services.InsertMessage(db, &message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	json.NewEncoder(w).Encode(map[string]interface{} {
		"message" : "Mensagem enviada com sucesso",
		"mensagem" : &message,
	})
}

func HandlerGetMessage (w http.ResponseWriter, r *http.Request, db *sql.DB) {
	message, err := services.GetMessagens(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	json.NewEncoder(w).Encode(map[string]interface{} {
		"message" : "listada com sucesso",
		"mensagem" : message,
	})
}