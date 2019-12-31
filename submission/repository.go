package submission

import "github.com/gustavonahum/code-validator/models"

type Repository interface {
	GetById(id int64) (*models.Submission, error)
	Store(submission *models.Submission) (*models.Submission, error)
	Delete(id int64) (*models.Submission, error)
}
