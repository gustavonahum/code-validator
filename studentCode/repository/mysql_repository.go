package repository

import (
	"database/sql"

	"github.com/gustavonahum/code-validator/studentCode"
)

type mysqlStudentCodeRepository struct {
	Conn *sql.DB
}

func NewMysqlStudentCodeRepository(Conn *sql.DB) studentCode.Repository {
	return &mysqlStudentCodeRepository{Conn}
}
