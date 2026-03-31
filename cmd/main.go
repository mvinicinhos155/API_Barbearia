package main

import (
	"api_barbearia/internal/database"
	"api_barbearia/internal/handlers"
	"api_barbearia/internal/middleware"
	"api_barbearia/internal/jobs"
	"log"
	"net/http"
	"github.com/rs/cors"
	"github.com/joho/godotenv"
)

func main() {


	err := godotenv.Load()
	if err != nil{log.Printf("Erro ao carregar o arquivo .env %s" , err)}


	db, err := database.Connect()
		if err != nil {
			log.Println("Erro ao conectar com bancos de dados", err)
		} 

	mux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type",  "Authorization"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(mux)

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			handler.HandlerCreateUser(w, r, db)
		}
	})

	mux.Handle("/users", middleware.AuthMiddleware(middleware.PermisionAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.HandlerGetUsers(w, r, db)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}))))

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.HandlerLogin(w, r, db)
		}
	})

		mux.Handle("/usersID", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodDelete:
			handler.HandlerDeleteUser(w, r, db)
		}
	})))

	mux.Handle("/haircut", middleware.AuthMiddleware(middleware.PermisionAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			handler.HandlerCreateHairs(w, r, db)
		}
	}))))

	mux.HandleFunc("/haircuts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.HandlerGetAllHairs(w, r, db)
		}
	})

	mux.HandleFunc("/haircutID", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodDelete:
			handler.HandlerDeleteHair(w, r, db)
		}
	})

	mux.Handle("/appointment", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			handler.HandlerCreateAppointment(w, r, db)
		}
	})))

	mux.Handle("/appoint", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.HandlerGetByUserId(w, r, db)
		}
	})))

	mux.Handle("/all_appointments", middleware.AuthMiddleware(middleware.PermisionAdmin(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.HandlerGetAllAppointment(w, r, db)
		}
	}))))

	mux.HandleFunc("/appointments", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handler.HandlerGetByDate(w, r, db)
		return
	}

	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
})

	mux.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			handler.HandlerSendMessage(w, r, db)
		}
	})

	mux.HandleFunc("/messagens", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			handler.HandlerGetMessage(w, r, db)
		}
	})


	jobs.StartReminderJob(db)

	database.Migrations(db)


	log.Println("Servidor rodando na porta 8080")
                                                                                      
	http.ListenAndServe(":8080", corsHandler)
}