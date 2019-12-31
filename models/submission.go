package models

import "time"

type Submission struct {
	Id        int64     `json:"id"`
	IdUser    int64     `json:"id_user"`
	IdCode    int64     `json:"id_code"`
	CreatedAt time.Time `json:"created_at"`
}
