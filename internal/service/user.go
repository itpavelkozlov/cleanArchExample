package service

import (
	"cleanArch/internal/models"
	"context"
)

type User interface {
	Create(ctx context.Context, user *models.User) error
	CreateAdmin(ctx context.Context, user *models.User) error
	CreateModerator(ctx context.Context, user *models.User) error

	Update(ctx context.Context, id string, user *models.User) error
	Find(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	Delete(ctx context.Context, id string) error

	AddUserToGroup(ctx context.Context, userId, groupId string)

	AdminList(ctx context.Context) ([]*models.User, error)
	ModeratorList(ctx context.Context) ([]*models.User, error)
	TodayBirthdayList(ctx context.Context) ([]*models.User, error)
}
