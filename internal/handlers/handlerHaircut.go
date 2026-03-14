package handler

import (
	"database/sql"
	"api_barbearia/internal/models"
	"api_barbearia/internal/services"
	"encoding/json"
	"net/http"
)

func HandlerCreateHaircut(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var hairCut models.Haircuts

	err := json.NewDecoder(r.Body).Decode(&hairCut)
		if err != nil {
			http.Error(w, "Erro ao enviar dados", http.StatusBadRequest)
			return
		}

	 err = services.InsertHairCut(db, &hairCut)
	   if err != nil {
		  http.Error(w, "Erro ao criar corte de cabelo", http.StatusBadRequest)
		  return
	   }

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Corte criado com sucesso",
		"Corte" : hairCut,
	})

}