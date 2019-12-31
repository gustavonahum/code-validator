package repository

import (
	"database/sql"

	"github.com/gustavonahum/code-validator/userCode"
)

type postgresqlUserCodeRepository struct {
	Conn *sql.DB
}

func NewPostgresqlUserCodeRepository(Conn *sql.DB) userCode.Repository {
	return &postgresqlUserCodeRepository{Conn}
}
