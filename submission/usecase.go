package submission

import "github.com/gustavonahum/code-validator/models"

type Usecase interface {
	GetById(idUser int64, idProblem int64) (*models.Submission, error)
	Store(submission *models.Submission) (*models.Submission, error)
	Delete(idUser int64, idProblem int64) (*models.Submission, error)
}
