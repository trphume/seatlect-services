package businessmb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
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

type Repo interface {
	ListBusiness(ctx context.Context, searchParams typedb.ListBusinessParams) ([]typedb.Business, error)
	ListBusinessByIds(ctx context.Context, ids []string) ([]typedb.Business, error)
}
