package orderwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/order_api"
	"net/http"
)

type Server struct {
	Repo Repo
}

func (s *Server) PostOrderVerify(ctx echo.Context) error {
	var req order_api.VerifyRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Could not bind request")
	}

	// Call Repo
	order, err := s.Repo.GetOrderWithReservationId(ctx.Request().Context(), req.OrderId, req.ReservationId)
	if err != nil {
		if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Invalid arguments")
		} else if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Could not find matching reservation and order")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	if err := s.Repo.UpdateOrderStatus(ctx.Request().Context(), req.OrderId, "USED"); err != nil {
		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	// Construct response
	seats := make([]string, len(order.Seats))
	for i := 0; i < len(order.Seats); i++ {
		seats[i] = order.Seats[i].Name
	}

	res := order_api.VerifyResponse{Seats: &seats}
	return ctx.JSONPretty(http.StatusOK, &res, "  ")
}

type Repo interface {
	GetOrderWithReservationId(ctx context.Context, orderId string, reservationId string) (*typedb.Order, error)
	UpdateOrderStatus(ctx context.Context, orderId string, status string) error
}
