package repository

import (
	"database/sql"

	"github.com/gustavonahum/code-validator/models"
	"github.com/gustavonahum/code-validator/user"
)

type postgresqlUserRepository struct {
	Conn *sql.DB
}

func NewPostgresqlUserRepository(Conn *sql.DB) user.Repository {
	return &postgresqlUserRepository{Conn}
}

func (u *postgresqlUserRepository) GetById(id int64) (*models.User, error) {
	return &models.User{}, nil
}
