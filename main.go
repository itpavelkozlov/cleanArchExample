package main

import (
	"cleanArch/internal/repository/mongo"
	"cleanArch/internal/resolver"
	"cleanArch/internal/server"
	"cleanArch/internal/service/service"
	"cleanArch/pkg/logger"
	"cleanArch/pkg/mongodb"
	"context"
	"os"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger()
	log.Info("Application starting")

	db, err := mongodb.NewMongoDB(ctx, log)
	if err != nil {
		log.Error("Can not start database")
		os.Exit(1)
	}
	log.Info("Mongodb started")
	userRepo := mongo.NewUserRepo(db)
	userService := service.NewUserService(userRepo, log)
	r := resolver.NewResolver(userService, log)
	router := server.NewRouter(r)
	httpServer := server.NewHttpServer(router)

	httpServer.ListenAndServe()
}
