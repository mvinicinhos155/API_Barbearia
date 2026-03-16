package handler

import (
	"api_barbearia/internal/models"
	"api_barbearia/internal/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerCreateHaircut(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var hairCut models.Haircuts

	userClaims := r.Context().Value("userID").(int)
	fmt.Println("USER ID:", userClaims)

	err := json.NewDecoder(r.Body).Decode(&hairCut)
		if err != nil {
			http.Error(w, "Erro ao enviar dados", http.StatusBadRequest)
			return
		}

		hair := models.Haircuts{
			Haircut_style: hairCut.Haircut_style,
			Price: hairCut.Price,
			Day: hairCut.Day,
			Hour: hairCut.Hour,
			User_id: userClaims,
		}

	 err = services.InsertHairCut(db, &hair)
	   if err != nil {
		  http.Error(w, "Erro ao criar corte de cabelo", http.StatusBadRequest)
		  return
	   }

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Corte criado com sucesso",
		"Corte" : hair,
	})

}