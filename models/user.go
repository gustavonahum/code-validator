package models

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Login string `json:"login"`
}
