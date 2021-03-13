package reservationmb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/commonpb"
	"github.com/tphume/seatlect-services/internal/genproto/reservationpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const iso8601 = "2006-01-02T15:04:05-0700"

type Server struct {
	Repo Repo

	reservationpb.UnimplementedReservationServiceServer
}

func (s *Server) ListReservation(ctx context.Context, req *reservationpb.ListReservationRequest) (*reservationpb.ListReservationResponse, error) {
	if len(req.Id) <= 0 {
		return nil, status.Error(codes.Unauthenticated, "ID is not valid")
	}

	start, err := time.Parse(iso8601, req.Start)
	if err != nil {
		start = time.Time{}
	}

	end, err := time.Parse(iso8601, req.End)
	if err != nil {
		end = time.Time{}
	}

	reservations, err := s.Repo.ListReservation(ctx, req.Id, start, end)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error")
	}

	return &reservationpb.ListReservationResponse{Reservation: typedbToCommonpb(reservations)}, nil
}

type Repo interface {
	ListReservation(ctx context.Context, id string, start time.Time, end time.Time) ([]typedb.Reservation, error)
}

// Helper function
func typedbToCommonpb(resv []typedb.Reservation) []*commonpb.Reservation {
	panic("implement me")
}

func seatsToCommonpb(seats []typedb.ReservationSeat) []*commonpb.ReservationSeat {
	panic("implement me")
}