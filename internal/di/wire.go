// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"kratos-calc/internal/dao"
	"kratos-calc/internal/server/grpc"
	"kratos-calc/internal/server/http"
	"kratos-calc/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
