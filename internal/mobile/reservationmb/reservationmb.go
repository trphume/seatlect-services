package reservationmb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
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

func (s *Server) ReserveSeats(ctx context.Context, req *reservationpb.ReserveSeatsRequest) (*reservationpb.ReserveSeatsResponse, error) {
	o, err := s.Repo.ReserveSeats(ctx, req.ResId, req.UserId, req.Name)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return nil, status.Error(codes.NotFound, "Could not find with given id")
		} else if err == commonErr.INVALID {
			return nil, status.Error(codes.InvalidArgument, "Bad argument format")
		} else if err == commonErr.CONFLICT {
			return nil, status.Error(codes.AlreadyExists, "Seat have already been reserved")
		}

		return nil, status.Error(codes.Internal, "Database error")
	}

	res := reservationpb.ReserveSeatsResponse{
		Order: &commonpb.Order{
			XId:           o.Id.Hex(),
			ReservationId: o.ReservationId.Hex(),
			Business:      o.BusinessId.Hex(),
			Start:         o.Start.Format(iso8601),
			End:           o.End.Format(iso8601),
			Seats:         orderSeatsToCommonpb(o.Seats),
			Status:        o.Status,
			Image:         o.Image,
			ExtraSpace:    int32(o.ExtraSpace),
			Name:          o.Name,
		},
	}

	return &res, nil
}

type Repo interface {
	ListReservation(ctx context.Context, id string, start time.Time, end time.Time) ([]typedb.Reservation, error)
	SearchReservation(ctx context.Context, searchParams typedb.SearchReservationParams) ([]typedb.Reservation, error)
	ReserveSeats(ctx context.Context, id string, user string, seats []string) (*typedb.Order, error)
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

func typedbOrderToCommonpb(orders []typedb.Order) []*commonpb.Order {
	res := make([]*commonpb.Order, len(orders))
	for i, o := range orders {
		res[i] = &commonpb.Order{
			XId:           o.Id.Hex(),
			ReservationId: o.ReservationId.Hex(),
			Business:      o.BusinessId.Hex(),
			Start:         o.Start.Format(iso8601),
			End:           o.End.Format(iso8601),
			Seats:         orderSeatsToCommonpb(o.Seats),
			Status:        o.Status,
			Image:         o.Image,
			ExtraSpace:    int32(o.ExtraSpace),
			Name:          o.Name,
		}
	}

	return res
}

func orderSeatsToCommonpb(seats []typedb.Seat) []*commonpb.OrderSeat {
	res := make([]*commonpb.OrderSeat, len(seats))
	for i, s := range seats {
		res[i] = &commonpb.OrderSeat{
			Name:  s.Name,
			Space: int32(s.Space),
		}
	}

	return res
}
