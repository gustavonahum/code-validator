package repository

import (
	"database/sql"

	"github.com/gustavonahum/code-validator/code"
)

type postgresqlCodeRepository struct {
	Conn *sql.DB
}

func NewPostgresqlCodeRepository(Conn *sql.DB) code.Repository {
	return &postgresqlCodeRepository{Conn}
}
