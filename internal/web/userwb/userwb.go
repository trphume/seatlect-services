package userwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/user_api"
	"net/http"
	"time"
)

type Server struct {
	Repo Repo
}

func (s *Server) PostUserLogin(ctx echo.Context) error {
	var req user_api.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	business := &typedb.Business{Username: req.Username, Password: req.Password}

	token, err := s.Repo.AuthenticateBusiness(ctx.Request().Context(), business)
	if err != nil {
		if err == commonErr.INTERNAL {
			return ctx.String(http.StatusInternalServerError, "Database error")
		}

		return ctx.String(http.StatusNotFound, "Username and password does not match")
	}

	// Construct response
	ctx.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	})

	res := user_api.LoginResponse{Id: createString(business.Id.Hex())}

	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) PostUserRegister(ctx echo.Context) error {
	var req user_api.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	business := &typedb.Business{
		Username:     *req.Username,
		Email:        *req.Email,
		Password:     *req.Password,
		BusinessName: *req.BusinessName,
		Type:         *req.Type,
		Tags:         make([]string, 0),
		Description:  *req.Description,
		Location: typedb.Location{
			Type:        "Point",
			Coordinates: []float64{float64(*req.Location.Longitude), float64(*req.Location.Latitude)},
		},
		Address:      *req.Address,
		DisplayImage: "",
		Images:       make([]string, 0),
		Placement:    make([]typedb.Seat, 0),
		Menu:         make([]typedb.MenuItems, 0),
		Status:       0,
		Verified:     false,
	}

	if err := s.Repo.CreateBusiness(ctx.Request().Context(), business); err != nil {
		if err == commonErr.INTERNAL {
			return ctx.String(http.StatusInternalServerError, "Database error")
		}

		return ctx.String(http.StatusConflict, "Business with that credentials already exist")
	}

	return ctx.String(http.StatusCreated, "Business created")
}

type Repo interface {
	// business is an out parameter - values will be overwritten
	// will return a user token
	AuthenticateBusiness(ctx context.Context, business *typedb.Business) (string, error)

	// business is an out parameter - values will be overwritten
	CreateBusiness(ctx context.Context, business *typedb.Business) error
}

// Helper function
func createString(s string) *string {
	return &s
}
