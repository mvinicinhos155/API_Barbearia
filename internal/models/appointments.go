package models

import "time"

type Appointments struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
	HaircutId int `json:"haircut_id"`
	Date   time.Time `json:"date"`
	Notified bool `json:"notified"`
}