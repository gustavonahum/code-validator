package models

import "time"

type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}