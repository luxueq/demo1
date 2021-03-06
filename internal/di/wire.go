// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"demo1/internal/dao"
	"demo1/internal/service"
	"demo1/internal/server/grpc"
	"demo1/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
