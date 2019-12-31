package usecase

import (
	"github.com/gustavonahum/code-validator/user"
	"github.com/gustavonahum/code-validator/userCode"
)

type userCodeUsecase struct {
	userCodeRepo userCode.Repository
	userRepo     user.Repository
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewUserCodeUsecase(uc userCode.Repository, u user.Repository) userCode.Usecase {
	return &userCodeUsecase{
		userCodeRepo: uc,
		userRepo:     u,
	}
}
