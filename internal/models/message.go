package models

type Message struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
}