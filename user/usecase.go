package user

import "github.com/gustavonahum/code-validator/models"

type Usecase interface {
	GetById(id int64) (*models.User, error)
	Store(user *models.User) (*models.User, error)
}
