package user

import "github.com/gustavonahum/code-validator/models"

type Repository interface {
	GetById(id int64) (*models.User, error)
}
