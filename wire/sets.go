package wire

import (
	"cleanArch/internal/repository/mongo"
	"cleanArch/internal/resolver"
	"cleanArch/internal/server"
	"cleanArch/internal/service/service"
	"cleanArch/pkg/logger"
	"cleanArch/pkg/mongodb"
	"github.com/google/wire"
)

var httpServerSet = wire.NewSet(
	server.NewRouter,
	server.NewHttpServer,
	resolver.NewResolver,
)

var serviceSet = wire.NewSet(
	service.NewUserService,
)

var repositorySet = wire.NewSet(
	mongo.NewUserRepo,
	//	postgres.NewUserRepo
)

var pkgSet = wire.NewSet(
	logger.NewLogger,
	mongodb.NewMongoDB,
)
