package handler

import (
	"api_barbearia/internal/models"
	"api_barbearia/internal/services"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandlerCreateHairs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var hair models.Haircuts

	err := json.NewDecoder(r.Body).Decode(&hair)
		if err != nil {
			http.Error(w, "Erro ao enviar dados", http.StatusBadRequest)
			return
		}

	err = services.InsertHaircut(db, &hair)
		if err != nil {
			http.Error(w, "Erro com banco de dados", http.StatusBadRequest)
			return
		}

	json.NewEncoder(w).Encode(map[string]interface{} {
		"message" : "Corte de cabelo criado",
		"haircut" : hair,
	})
}

func HandlerGetAllHairs(w http.ResponseWriter, r *http.Request,  db *sql.DB) {

	hairs, err := services.GetAllHairs(db)
		if err != nil {
			http.Error(w, "Erro com o banco de dados", http.StatusBadRequest)
			return
		}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Cortes de cabelos listado com sucesso",
		"haircuts" : hairs,
	})
}

func HandlerDeleteHair(w http.ResponseWriter,  r *http.Request, db *sql.DB) {
	var hairId models.Haircuts

	err := json.NewDecoder(r.Body).Decode(&hairId)
		if err != nil {
			http.Error(w, "Erro ao enviar dados", http.StatusBadRequest)
			return
		}

	err = services.DeleteHair(db, hairId.ID)
		if err != nil {
			http.Error(w, "Erro ao excluir o corte de cabelo", http.StatusBadRequest)
			return
		}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Corte de cabelo deletado com sucesso",
	})
}