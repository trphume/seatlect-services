package orderwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/typedb"
)

type Server struct {
	Repo Repo
}

func (s *Server) PostOrderVerify(ctx echo.Context) error {
	panic("implement me")
}

type Repo interface {
	GetOrderWithReservationId(ctx context.Context, orderId string, reservationId string) (*typedb.Order, error)
	UpdateOrderStatus(ctx context.Context, orderId string, status string) error
}
