package reservationmb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/commonpb"
	"github.com/tphume/seatlect-services/internal/genproto/reservationpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *Server) ReserveSeats(context.Context, *reservationpb.ReserveSeatsRequest) (*reservationpb.ReserveSeatsResponse, error) {
	panic("implement me")
}

type Repo interface {
	ListReservation(ctx context.Context, id string, start time.Time, end time.Time) ([]typedb.Reservation, error)
}

// Helper function
func typedbToCommonpb(resv []typedb.Reservation) []*commonpb.Reservation {
	res := make([]*commonpb.Reservation, len(resv))
	for i, r := range resv {
		res[i] = &commonpb.Reservation{
			Id:         r.Id.Hex(),
			BusinessId: r.BusinessId.Hex(),
			Name:       r.Name,
			Start:      r.Start.Format(iso8601),
			End:        r.End.Format(iso8601),
			Placement: &commonpb.ReservationPlacement{
				Height: int32(r.Placement.Height),
				Width:  int32(r.Placement.Width),
				Seats:  seatsToCommonpb(r.Placement.Seats),
			},
			Image: r.Image,
		}
	}

	return res
}

func seatsToCommonpb(seats []typedb.ReservationSeat) []*commonpb.ReservationSeat {
	res := make([]*commonpb.ReservationSeat, len(seats))
	for i, s := range seats {
		var userId string
		if primitive.ObjectID.IsZero(s.User) {
			userId = ""
		} else {
			userId = s.User.Hex()
		}

		res[i] = &commonpb.ReservationSeat{
			Name:     s.Name,
			Floor:    int32(s.Floor),
			Type:     s.Type,
			Space:    int32(s.Space),
			User:     userId,
			Status:   s.Status,
			X:        s.X,
			Y:        s.Y,
			Width:    s.Width,
			Height:   s.Height,
			Rotation: s.Rotation,
		}
	}

	return res
}
