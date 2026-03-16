package handler

import (
	"api_barbearia/internal/models"
	"api_barbearia/internal/services"
	"api_barbearia/internal/token"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"golang.org/x/crypto/bcrypt"
)

func HandlerCreateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var user models.Users

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao enviar dados", http.StatusBadRequest)
		return
	}

	var hasUpper = regexp.MustCompile(`[A-Z]`)
	var hasNumber = regexp.MustCompile(`[0-9]`)

	if len(user.Password) < 8 || !hasUpper.MatchString(user.Password) || !hasNumber.MatchString(user.Password) {
		http.Error(w, "Senha muito curta, tente novamente", http.StatusUnauthorized)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)


	err = services.InsertUser(db, &user)
		if err != nil {
			fmt.Println("Erro real:", err)
			http.Error(w, "Erro ao criar usuário", http.StatusBadRequest)
			return
		}

	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func HandlerLogin (w http.ResponseWriter, r *http.Request, db *sql.DB){

	var Login struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	 err := json.NewDecoder(r.Body).Decode(&Login)
	 	if err != nil {
			http.Error(w, "Erro ao enviar dados", http.StatusBadRequest)
			return 
		}

	dbUser, err := services.GetUserbyEmail(db, Login.Email)
		if err != nil {
			http.Error(w, "Erro ao fazer login", http.StatusBadRequest)
			return 	
		}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(Login.Password))
		if err != nil {
			http.Error(w, "Email ou senha invalidos", http.StatusUnauthorized)
			return 
		} 
		
	tokenString, err := token.GenerateJwt(&dbUser)
			if err != nil {
				http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message" : "login feito com sucesso",
				"login" : dbUser,
				"token" : tokenString,
			})

}

func HandlerGetUsers (w http.ResponseWriter, r *http.Request, db *sql.DB) {

	AllUsers, err := services.GetAllUser(db)
		if err !=  nil {
			http.Error(w,"Erro ao  pegar todos os usuários", http.StatusInternalServerError)
			return
		}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Usuários listados com sucesso",
		"Usuários" : AllUsers,
	})
}