package repository

import (
	"cleanArch/internal/models"
	"context"
)

type User interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, id string, user *models.User) error
	Find(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	Delete(ctx context.Context, id string) error
}
