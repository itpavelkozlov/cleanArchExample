//+build wireinject

package wire

import (
	"context"
	"github.com/google/wire"
)

func InitializeApplication(ctx context.Context) (Application, error) {
	wire.Build(
		newApplication,
		httpServerSet,
		serviceSet,
		repositorySet,
		pkgSet,
	)
	return Application{}, nil
}
