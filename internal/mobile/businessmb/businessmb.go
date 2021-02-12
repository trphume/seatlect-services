package businessmb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/genproto/businesspb"
)

type Server struct {
	businesspb.UnimplementedBusinessServiceServer
}

func (s *Server) ListBusiness(ctx context.Context, request *businesspb.ListBusinessRequest) (*businesspb.ListBusinessResponse, error) {
	panic("implement me")
}

func (s *Server) ListBusinessById(ctx context.Context, request *businesspb.ListBusinessByIdRequest) (*businesspb.ListBusinessByIdResponse, error) {
	panic("implement me")
}
