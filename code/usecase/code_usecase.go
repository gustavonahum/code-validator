package usecase

import (
	"github.com/gustavonahum/code-validator/code"
	"github.com/gustavonahum/code-validator/user"
)

type codeUsecase struct {
	codeRepo code.Repository
	userRepo user.Repository
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewCodeUsecase(c code.Repository, u user.Repository) code.Usecase {
	return &codeUsecase{
		codeRepo: c,
		userRepo: u,
	}
}
