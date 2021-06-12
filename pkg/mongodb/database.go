package mongodb

import (
	"cleanArch/pkg/logger"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(ctx context.Context, logger logger.Logger) (*mongo.Database, error) {
	dbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?MaxPoolSize=50&MinPoolSize=20", "admin", "admin", "127.0.0.1", "27017")
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	logger.Info("Database is connected.")
	return mongoClient.Database("prod"), nil
}
