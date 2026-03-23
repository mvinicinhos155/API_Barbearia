package main

import (
	"api_barbearia/internal/database"
	"api_barbearia/internal/handlers"
	"api_barbearia/internal/middleware"
	"api_barbearia/internal/jobs"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {


	err := godotenv.Load()
	if err != nil{log.Printf("Erro ao carregar o arquivo .env %s" , err)}


	db, err := database.Connect()
		if err != nil {
			log.Println("Erro ao conectar com bancos de dados", err)
		} 

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			handler.HandlerCreateUser(w, r, db)
		}
	})

	http.Handle("/users", middleware.AuthMiddleware(middleware.PermisionAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.HandlerGetUsers(w, r, db)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}))))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.HandlerLogin(w, r, db)
		}
	})

		http.Handle("/usersID", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodDelete:
			handler.HandlerDeleteUser(w, r, db)
		}
	})))

	http.Handle("/haircut", middleware.AuthMiddleware(middleware.PermisionAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			handler.HandlerCreateHairs(w, r, db)
		}
	}))))

	http.HandleFunc("/haircuts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.HandlerGetAllHairs(w, r, db)
		}
	})

	http.HandleFunc("/haircutID", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodDelete:
			handler.HandlerDeleteHair(w, r, db)
		}
	})

	http.Handle("/appointment", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			handler.HandlerCreateAppointment(w, r, db)
		}
	})))

	http.Handle("/appoint", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.HandlerGetByUserId(w, r, db)
		}
	})))


	jobs.StartReminderJob(db)

	database.Migrations(db)


	log.Println("Servidor rodando na porta 8080")
                                                                                      
	http.ListenAndServe(":8080", nil)
}