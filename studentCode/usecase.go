package studentCode

import (
	"context"
	"models"
)

type Usecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*models.StudentCode, string, error)
	GetByID(ctx context.Context, id int64) (*models.StudentCode, error)
	Update(ctx context.Context, ar *models.StudentCode) error
	GetByName(ctx context.Context, name string) (*models.StudentCode, error)
	Store(context.Context, *models.StudentCode) error
	Delete(ctx context.Context, id int64) error
}