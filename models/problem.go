package models

import "time"

type Problem struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	TimeLimit   time.Time `json:"time_limit"`
	IdTestCases int64     `json:"id_test_cases"`
}
