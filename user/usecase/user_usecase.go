package usecase

import (
	"github.com/gustavonahum/code-validator/models"
	"github.com/gustavonahum/code-validator/user"
)

type userUsecase struct {
	userRepo user.Repository
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewUserUsecase(u user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) GetById(id int64) (*models.User, error) {
	return &models.User{}, nil
}
