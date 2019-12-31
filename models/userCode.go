package models

import "time"

type UserCode struct {
	Id int64 `json:"id"`
	Filename string `json:"filename"`
	Language string `json:"language"`
	LanguageVersion string `json:"language_version"`
	Code string `json:"code"`
	IdUser int64 `json:"id_user"`
}