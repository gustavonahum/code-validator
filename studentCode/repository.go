package studentCode

type Repository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []*models.StudentCode, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (*models.StudentCode, error)
	GetByTitle(ctx context.Context, title string) (*models.StudentCode, error)
	Update(ctx context.Context, ar *models.StudentCode) error
	Store(ctx context.Context, a *models.StudentCode) error
	Delete(ctx context.Context, id int64) error
}