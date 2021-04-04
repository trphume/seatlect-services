package placementwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/placement_api"
	"net/http"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetPlacementBusinessId(ctx echo.Context, businessId string) error {
	placement, err := s.Repo.GetPlacement(ctx.Request().Context(), businessId)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Cannot find placement of given business id")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Business id format is incorrect")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := typedbToOapi(*placement)
	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) PutPlacementBusinessId(ctx echo.Context, businessId string) error {
	var req placement_api.PutPlacementBusinessIdJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	placement := oapiToTypedb(placement_api.Placement(req))

	if err := s.Repo.UpdatePlacement(ctx.Request().Context(), businessId, placement); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Cannot find placement of given business id")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Business id format is incorrect")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	return ctx.String(http.StatusOK, "Placement updated successfully")
}

type Repo interface {
	GetPlacement(ctx context.Context, id string) (*typedb.Placement, error)
	UpdatePlacement(ctx context.Context, id string, placement typedb.Placement) error
}

// Parsing function
func typedbToOapi(placement typedb.Placement) placement_api.Placement {
	return placement_api.Placement{
		Height: &placement.Height,
		Width:  &placement.Width,
		Seats:  typedbSeatsToOapi(placement.Seats),
	}
}

func typedbSeatsToOapi(seats []typedb.Seat) *[]placement_api.Seat {
	res := make([]placement_api.Seat, len(seats))
	for i, s := range seats {
		res[i] = placement_api.Seat{
			Floor:    &s.Floor,
			Height:   createFloat32(float32(s.Height)),
			Name:     &s.Name,
			Rotation: createFloat32(float32(s.Rotation)),
			Space:    &s.Space,
			Y:        createFloat32(float32(s.Y)),
			Width:    createFloat32(float32(s.Width)),
			X:        createFloat32(float32(s.X)),
			Type:     &s.Type,
		}
	}

	return &res
}

func oapiToTypedb(placement placement_api.Placement) typedb.Placement {
	return typedb.Placement{
		Width:  *placement.Width,
		Height: *placement.Height,
		Seats:  oapiSeatsToTypedb(*placement.Seats),
	}
}

func oapiSeatsToTypedb(seats []placement_api.Seat) []typedb.Seat {
	res := make([]typedb.Seat, len(seats))
	for i, s := range seats {
		res[i] = typedb.Seat{
			Name:     *s.Name,
			Floor:    *s.Floor,
			Type:     *s.Type,
			Space:    *s.Space,
			X:        float64(*s.X),
			Y:        float64(*s.Y),
			Width:    float64(*s.Width),
			Height:   float64(*s.Height),
			Rotation: float64(*s.Rotation),
		}
	}

	return res
}

// Helper function
func createString(s string) *string {
	return &s
}

func createFloat32(f float32) *float32 {
	return &f
}
