package models

type Haircuts struct {
	ID int `json:"id"`
	Haircut_style string `json:"haircut_style"`
	Price float64 `json:"price"`
	Day string `json:"day"`
	Hour string `json:"hour"`
	User_id int `json:"user_id"`
}