package reservationwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"time"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetReservationBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PostReservationBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

type Repo interface {
	ListReservation(ctx context.Context, id string, start time.Time, end time.Time) ([]typedb.Reservation, error)
	CreateReservation(ctx context.Context, placement typedb.Placement) error
}
