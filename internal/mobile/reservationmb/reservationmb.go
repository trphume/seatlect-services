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
	"sync"
	"time"
)

const iso8601 = "2006-01-02T15:04:05-0700"

type Server struct {
	mu sync.Mutex

	Repo               Repo
	SubscribersChannel map[string]map[chan typedb.ReservationPlacement]bool

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

func (s *Server) SearchReservation(ctx context.Context, req *reservationpb.SearchReservationRequest) (*reservationpb.SearchReservationResponse, error) {
	// Construct params
	start, err := time.Parse(iso8601, req.Start)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Start time format does not match")
	}

	end, err := time.Parse(iso8601, req.End)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "End time format does not match")
	}

	if start.After(end) {
		return nil, status.Error(codes.InvalidArgument, "End time cannot be before start time")
	}

	searchParams := typedb.SearchReservationParams{
		Name: req.Name,
		Type: req.Type,
		Location: typedb.Location{
			Type:        "Point",
			Coordinates: []float64{req.Location.Longitude, req.Location.Latitude},
		},
		Start: start,
		End:   end,
	}

	// Call Repo
	reservations, err := s.Repo.SearchReservation(ctx, searchParams)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error")
	}

	return &reservationpb.SearchReservationResponse{Reservation: typedbToCommonpb(reservations)}, nil
}

func (s *Server) ReserveSeats(ctx context.Context, req *reservationpb.ReserveSeatsRequest) (*reservationpb.ReserveSeatsResponse, error) {
	// attempt to make reservations
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

	// notifity subscribers of changes - indepdendent of this handler
	go s.notifySubscribers(req.ResId)

	return &res, nil
}

func (s *Server) Susbscribe(req *reservationpb.SubscribeRequest, serv reservationpb.ReservationService_SusbscribeServer) error {
	// fetch initial data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	initial, err := s.Repo.GetReservationById(ctx, "", req.Id)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return status.Error(codes.NotFound, "reservation not found with given id")
		} else if err == commonErr.INVALID {
			return status.Error(codes.InvalidArgument, "error not found")
		}

		return status.Error(codes.Internal, "database error")
	}

	res := reservationpb.SubscribeReponse{Placement: &commonpb.ReservationPlacement{
		Width:  int32(initial.Placement.Width),
		Height: int32(initial.Placement.Height),
		Seats:  seatsToCommonpb(initial.Placement.Seats),
	}}

	if err := serv.Send(&res); err != nil {
		return status.Error(codes.Unknown, "connection failed")
	}

	// new channel to receive from
	c := make(chan typedb.ReservationPlacement)
	defer close(c)

	// update reservation subscribers channel map
	s.mu.Lock()
	if _, ok := s.SubscribersChannel[req.Id]; !ok {
		// initialize an array if does not exist
		s.SubscribersChannel[req.Id] = make(map[chan typedb.ReservationPlacement]bool)
	}

	s.SubscribersChannel[req.Id][c] = true

	s.mu.Unlock()

	// wait for notifications and send updates to client
	for {
		placement := <-c

		res := reservationpb.SubscribeReponse{Placement: &commonpb.ReservationPlacement{
			Width:  int32(placement.Width),
			Height: int32(placement.Height),
			Seats:  seatsToCommonpb(placement.Seats),
		}}

		if err := serv.Send(&res); err != nil {
			// remove subscriber channel
			s.mu.Lock()
			delete(s.SubscribersChannel[req.Id], c)
			s.mu.Unlock()

			return status.Error(codes.Unknown, "connection failed")
		}
	}
}

// helper functions
func (s *Server) notifySubscribers(resId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resv, err := s.Repo.GetReservationById(ctx, "", resId)
	if err == nil {
		if subs, ok := s.SubscribersChannel[resId]; ok {
			for k, _ := range subs {
				k <- resv.Placement
			}
		}
	}
}

type Repo interface {
	ListReservation(ctx context.Context, id string, start time.Time, end time.Time) ([]typedb.Reservation, error)
	SearchReservation(ctx context.Context, searchParams typedb.SearchReservationParams) ([]typedb.Reservation, error)
	ReserveSeats(ctx context.Context, id string, user string, seats []string) (*typedb.Order, error)
	GetReservationById(ctx context.Context, businessId string, reservationId string) (*typedb.Reservation, error)
}

// Parsing function
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
