package adminwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/gen_openapi/admin_api"
	"net/http"
	"time"
)

type Server struct {
	Repo Repo
}

func (s *Server) PostAdminLogin(ctx echo.Context) error {
	var req admin_api.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusInternalServerError, "Error parsing request body")
	}

	token, err := s.Repo.AuthenticateAdmin(ctx.Request().Context(), req.Username, req.Password)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusUnauthorized, "Credentials does not match")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	ctx.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return ctx.String(http.StatusNoContent, "Authentication was successful")
}

type Repo interface {
	AuthenticateAdmin(ctx context.Context, username string, password string) (string, error)
}
