package usecase

import (
	"github.com/gustavonahum/code-validator/models"
	"github.com/gustavonahum/code-validator/submission"
)

type submissionUsecase struct {
	submissionRepo submission.Repository
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewSubmissionUsecase(s submission.Repository) submission.Usecase {
	return &submissionUsecase{
		submissionRepo: s,
	}
}

func (s *submissionUsecase) GetById(idUser int64, idProblem int64) (*models.Submission, error) {
	return s.submissionRepo.GetById(idUser, idProblem)
}

func (s *submissionUsecase) Store(submission *models.Submission) (*models.Submission, error) {
	return s.submissionRepo.Store(submission)
}

func (s *submissionUsecase) Delete(idUser int64, idProblem int64) (*models.Submission, error) {
	return s.submissionRepo.Delete(idUser, idProblem)
}
