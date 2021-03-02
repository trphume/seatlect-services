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
		return ctx.JSONPretty(http.StatusBadRequest, "Error binding request body", "  ")
	}

	business := &typedb.Business{Username: req.Username, Password: req.Password}

	token, err := s.Repo.AuthenticateBusiness(ctx.Request().Context(), business)
	if err != nil {
		if err == commonErr.INTERNAL {
			return ctx.JSONPretty(http.StatusInternalServerError, "Database error", "  ")
		}

		return ctx.JSONPretty(http.StatusNotFound, "Username and password does not match", "  ")
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
	panic("implement me")
}

type Repo interface {
	// business is an out parameter - values will be overwritten
	// will return a user token
	AuthenticateBusiness(ctx context.Context, business *typedb.Business) (string, error)

	// business is an out parameter - values will be overwritten
	// will return a user token
	CreateBusiness(ctx context.Context, business *typedb.Business) (string, error)
}

// Helper function
func createString(s string) *string {
	return &s
}
