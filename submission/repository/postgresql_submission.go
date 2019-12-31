package repository

import (
	"database/sql"

	"github.com/gustavonahum/code-validator/models"
	"github.com/gustavonahum/code-validator/submission"
)

type postgresqlSubmissionRepository struct {
	Conn *sql.DB
}

func NewPostgresqlSubmissionRepository(Conn *sql.DB) submission.Repository {
	return &postgresqlSubmissionRepository{Conn}
}

func (u *postgresqlSubmissionRepository) GetById(idUser int64, idProblem int64) (*models.Submission, error) {
	return &models.Submission{}, nil
}

func (u *postgresqlSubmissionRepository) Store(user *models.Submission) (*models.Submission, error) {
	return user, nil
}

func (u *postgresqlSubmissionRepository) Delete(idUser int64, idProblem int64) (*models.Submission, error) {
	return &models.Submission{}, nil
}
