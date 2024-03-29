package reservationwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/reservation_api"
	"github.com/tphume/seatlect-services/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
	"net/http"
	"time"
)

const iso8601 = "2006-01-02T15:04:05-0700"

type Server struct {
	Repo     Repo
	UserRepo UserRepo

	Mail *gomail.Dialer
}

func (s *Server) GetReservationBusinessId(ctx echo.Context, businessId string, params reservation_api.GetReservationBusinessIdParams) error {
	start, err := time.Parse(iso8601, params.Start)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Error parsing start date")
	}

	end, err := time.Parse(iso8601, params.End)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Error parsing end date")
	}

	// Check that start is before end
	if start.After(end) {
		return ctx.String(http.StatusBadRequest, "Start time cannot be after end time")
	}

	reservations, err := s.Repo.ListReservation(ctx.Request().Context(), businessId, start, end)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Cannot find reservations of given business id")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Business id format is incorrect")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := reservation_api.ListReservationResponse{Reservations: typedbListToOapi(reservations)}
	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) PostReservationBusinessId(ctx echo.Context, businessId string) error {
	var req reservation_api.CreateReservationRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	start, err := time.Parse(iso8601, *req.Start)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Error parsing start date")
	}

	end, err := time.Parse(iso8601, *req.End)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Error parsing end date")
	}

	// Check that start is before end
	if start.After(end) {
		return ctx.String(http.StatusBadRequest, "Start time cannot be after end time")
	} else if start.Before(time.Now()) {
		return ctx.String(http.StatusBadRequest, "Start time cannot be before current time")
	}

	// Create empty Placement object
	id := primitive.NewObjectIDFromTimestamp(time.Now())

	bId, err := primitive.ObjectIDFromHex(businessId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Business id format is incorrect")
	}

	// Placement and image to be filled after querying business via repo
	reservation := typedb.Reservation{
		Id:         id,
		BusinessId: bId,
		Name:       *req.Name,
		Start:      start,
		End:        end,
		Placement:  typedb.ReservationPlacement{},
		Image:      "",
		Status:     1,
	}

	if err := s.Repo.CreateReservation(ctx.Request().Context(), &reservation); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Cannot find reservations of given business id")
		} else if err == commonErr.CONFLICT {
			return ctx.String(http.StatusConflict, "No overlapping resevation are allowed")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	resv := typedbToOapi(reservation)
	res := reservation_api.CreateReservationResponse{Reservation: &resv}

	return ctx.JSONPretty(http.StatusCreated, res, "  ")
}

func (s *Server) GetReservationBusinessIdReservationId(ctx echo.Context, businessId string, reservationId string) error {
	if businessId == "" {
		return ctx.String(http.StatusBadRequest, "missing businessid")
	}

	reservation, err := s.Repo.GetReservationById(ctx.Request().Context(), businessId, reservationId)
	if err != nil {
		if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "invalid id format")
		} else if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "reservation not found with given id")
		}

		return ctx.String(http.StatusNotFound, "database error")
	}

	resv := typedbToOapi(*reservation)
	res := reservation_api.GetReservationResponse{Reservation: &resv}

	return ctx.JSONPretty(http.StatusOK, &res, "  ")
}

func (s *Server) PatchReservationReservationIdStatus(ctx echo.Context, reservationId string) error {
	// update reservation status
	users, err := s.Repo.UpdateReservationStatus(ctx.Request().Context(), reservationId)
	if err != nil {
		if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Invalid id format")
		} else if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Reservation not found with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	// update order status
	go s.updateOrderStatus(reservationId)

	// get user email send email notification to each user
	go s.notifyReservationStatusUpdate(users)

	return ctx.String(http.StatusOK, "Updated successfully")
}

func (s *Server) updateOrderStatus(reservationId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_ = s.Repo.UpdateOrderStatus(ctx, reservationId)
}

func (s *Server) notifyReservationStatusUpdate(users []primitive.ObjectID) {
	// get list of emails to send
	ctx := context.Background()
	emails, _ := s.UserRepo.ListUserEmailById(ctx, users)

	// send email to each user
	for _, e := range emails {
		go utils.SendEmail(
			s.Mail,
			e,
			"Business have cancelled Reservation",
			"Unfortunately, one of your reservation have been cancelled by the Business. We apologize of any inconveniences.")
	}
}

type Repo interface {
	ListReservation(ctx context.Context, id string, start time.Time, end time.Time) ([]typedb.Reservation, error)
	CreateReservation(ctx context.Context, reservation *typedb.Reservation) error
	GetReservationById(ctx context.Context, businessId string, reservationId string) (*typedb.Reservation, error)
	UpdateReservationStatus(ctx context.Context, reservationId string) ([]primitive.ObjectID, error)
	UpdateOrderStatus(ctx context.Context, reservationId string) error
}

type UserRepo interface {
	ListUserEmailById(ctx context.Context, users []primitive.ObjectID) ([]string, error)
}

// Parsing function - kill me please
func typedbListToOapi(reservations []typedb.Reservation) *[]reservation_api.Reservation {
	res := make([]reservation_api.Reservation, len(reservations))
	for i, r := range reservations {
		res[i] = typedbToOapi(r)
	}

	return &res
}

func typedbToOapi(reservation typedb.Reservation) reservation_api.Reservation {
	return reservation_api.Reservation{
		End:       createString(reservation.End.Format(iso8601)),
		Name:      &reservation.Name,
		Placement: typedbPmtToOapi(reservation.Placement),
		Start:     createString(reservation.Start.Format(iso8601)),
		Id:        createString(reservation.Id.Hex()),
	}
}

func typedbPmtToOapi(placement typedb.ReservationPlacement) *reservation_api.Placement {
	return &reservation_api.Placement{
		Height: &placement.Height,
		Seats:  typedbSeatsToOapi(placement.Seats),
		Width:  &placement.Width,
	}
}

func typedbSeatsToOapi(seats []typedb.ReservationSeat) *[]reservation_api.Seat {
	res := make([]reservation_api.Seat, len(seats))
	for i, s := range seats {
		var userId string
		if primitive.ObjectID.IsZero(s.User) {
			userId = ""
		} else {
			userId = s.User.Hex()
		}

		res[i] = reservation_api.Seat{
			Floor:    createInt(s.Floor),
			Height:   createFloat32(float32(s.Height)),
			Name:     createString(s.Name),
			Rotation: createFloat32(float32(s.Rotation)),
			Space:    createInt(s.Space),
			Status:   createString(s.Status),
			Type:     createString(s.Type),
			User:     createString(userId),
			Username: createString(s.Username),
			Width:    createFloat32(float32(s.Width)),
			X:        createFloat32(float32(s.X)),
			Y:        createFloat32(float32(s.Y)),
		}
	}

	return &res
}

// Helper function
func createString(s string) *string {
	return &s
}

func createFloat32(f float32) *float32 {
	return &f
}

func createInt(i int) *int {
	return &i
}
