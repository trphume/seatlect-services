package reservationmb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/genproto/reservationpb"
	"time"
)

type Server struct {
	Repo Repo

	reservationpb.UnimplementedReservationServiceServer
}

func (s Server) ListReservation(ctx context.Context, request *reservationpb.ListReservationRequest) (*reservationpb.ListReservationResponse, error) {
	panic("implement me")
}

type Repo interface {
	ListReservation(ctx context.Context, id string, start time.Time, end time.Time)
}
