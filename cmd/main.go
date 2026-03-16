package main

import (
	"api_barbearia/internal/database"
	"api_barbearia/internal/handlers"
	"api_barbearia/internal/middleware"
	//"fmt"
	"log"
	"net/http"
	//"time"
	//"github.com/go-co-op/gocron"
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

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.HandlerLogin(w, r, db)
		}
	})

	http.Handle("/haircut", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			handler.HandlerCreateHaircut(w, r, db)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})))

	/*s := gocron.NewScheduler(time.Local)

	task := func ()  {
		fmt.Println("Testando o gocron, Olá...")
	}

	s.Every(5).Seconds().Do(task)
	s.Every(1).Friday().At(corte.horar).Do(func ()  {
		fmt.Println("Está na hora...")
	})

	go func ()  {
		s.StartAsync()
	}()

	select{}*/



	database.Migrations(db)


	log.Println("Servidor rodando na porta 8080")
                                                                                      
	http.ListenAndServe(":8080", nil)
}