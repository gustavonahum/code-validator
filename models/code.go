package models

type Code struct {
	Id       int64  `json:"id"`
	Filename string `json:"filename"`
	Code     string `json:"code"`
	Language string `json:"language"`
}
