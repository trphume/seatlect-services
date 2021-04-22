package orderwb

import (
	"context"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Repo Repo
}

func (s *Server) PostOrderVerify(ctx echo.Context) error {
	panic("implement me")
}

type Repo interface {
	CheckOrderReservationId(ctx context.Context, orderId string, reservationId string) error
	UpdateOrderStatus(ctx context.Context, orderId string, status string) error
}
