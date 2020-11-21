// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/taadis/kratos-calc/internal/dao"
	"github.com/taadis/kratos-calc/internal/server/grpc"
	"github.com/taadis/kratos-calc/internal/server/http"
	"github.com/taadis/kratos-calc/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
