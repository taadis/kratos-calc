package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	pb "github.com/taadis/kratos-calc/api"
	"github.com/taadis/kratos-calc/internal/dao"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.CalcServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Add
func (s *Service) Add(ctx context.Context, req *pb.AddRequest) (res *pb.AddResponse, err error) {
	res = &pb.AddResponse{
		C: req.A + req.B,
	}
	fmt.Println("Add", req.A, req.B, res.C)
	return
}

// Subtract
func (s *Service) Subtract(ctx context.Context, req *pb.SubtractRequest) (res *pb.SubtractResponse, err error) {
	res = &pb.SubtractResponse{
		C: req.A - req.B,
	}
	fmt.Println("Subtract:", req.A, "-", req.B, "=", res.C)
	return
}

// Close close the resource.
func (s *Service) Close() {
}
