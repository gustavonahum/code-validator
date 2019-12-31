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

func (s *submissionUsecase) GetById(id int64) (*models.Submission, error) {
	return s.submissionRepo.GetById(id)
}

func (s *submissionUsecase) Store(submission *models.Submission) (*models.Submission, error) {
	return s.submissionRepo.Store(submission)
}

func (s *submissionUsecase) Delete(id int64) (*models.Submission, error) {
	return s.submissionRepo.Delete(id)
}
