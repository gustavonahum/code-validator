package models

import (
	"time"

	"github.com/gustavonahum/code-validator/models/enum"
)

type Profile struct {
	Id          int64            `json:"id"`
	IdUser      int64            `json:"id_user"`
	ProfileType enum.ProfileType `json:"type"`
	CreatedAt   time.Time        `json:"created_at"`
}
