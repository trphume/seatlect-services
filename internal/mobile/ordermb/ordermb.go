package ordermb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/orderpb"
)

type Server struct {
	orderpb.UnimplementedOrderServiceServer
}

func (s *Server) ListOrder(ctx context.Context, request *orderpb.ListOrderRequest) (*orderpb.ListOrderResponse, error) {
	panic("implement me")
}

type Repo interface {
	ListOrderByCustomer(ctx context.Context, CustomerId string, limit int32, page int32) ([]typedb.Order, error)
}
