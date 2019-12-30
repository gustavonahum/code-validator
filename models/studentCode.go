package models

import "time"

type StudentCode struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Student Student `json:"student"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}