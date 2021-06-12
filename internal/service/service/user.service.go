package service

import (
	"cleanArch/internal/models"
	"cleanArch/internal/repository"
	"cleanArch/internal/service"
	"cleanArch/pkg/logger"
	"context"
	"errors"
)

type UserService struct {
	userRepo repository.User
	log      logger.Logger
}

func NewUserService(userRepo repository.User, log logger.Logger) service.User {
	return &UserService{
		userRepo: userRepo,
		log:      log,
	}
}

func (u UserService) Create(ctx context.Context, user *models.User) error {
	if user.Name != "Артем" {
		err := u.userRepo.Create(ctx, user)
		if err != nil {
			u.log.Error("Error happened!")
			return errors.New("internal err")
		}
		return err
	}
	return errors.New("invalid name")
}

func (u UserService) CreateAdmin(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (u UserService) CreateModerator(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (u UserService) Update(ctx context.Context, id string, user *models.User) error {
	panic("implement me")
}

func (u UserService) Find(ctx context.Context, id string) (*models.User, error) {
	panic("implement me")
}

func (u UserService) List(ctx context.Context) ([]*models.User, error) {
	panic("implement me")
}

func (u UserService) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (u UserService) AddUserToGroup(ctx context.Context, userId, groupId string) {
	panic("implement me")
}

func (u UserService) AdminList(ctx context.Context) ([]*models.User, error) {
	panic("implement me")
}

func (u UserService) ModeratorList(ctx context.Context) ([]*models.User, error) {
	panic("implement me")
}

func (u UserService) TodayBirthdayList(ctx context.Context) ([]*models.User, error) {
	panic("implement me")
}
