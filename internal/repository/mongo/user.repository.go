package mongo

import (
	"cleanArch/internal/models"
	"cleanArch/internal/repository"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewUserRepo(db *mongo.Database) repository.User {
	col := db.Collection("user")
	return &UserRepo{
		db:  db,
		col: col,
	}
}

func (u UserRepo) Create(ctx context.Context, user *models.User) error {
	_, err := u.col.InsertOne(ctx, user, nil)
	return err
}

func (u UserRepo) Update(ctx context.Context, id string, user *models.User) error {
	panic("implement me")
}

func (u UserRepo) Find(ctx context.Context, id string) (*models.User, error) {
	panic("implement me")
}

func (u UserRepo) List(ctx context.Context) ([]*models.User, error) {
	panic("implement me")
}

func (u UserRepo) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
