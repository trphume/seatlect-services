package ordermb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/genproto/orderpb"
)

type Server struct {
	orderpb.UnimplementedOrderServiceServer
}

func (s *Server) ListOrder(ctx context.Context, request *orderpb.ListOrderRequest) (*orderpb.ListOrderResponse, error) {
	panic("implement me")
}
