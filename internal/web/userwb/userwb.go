package userwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/typedb"
)

type Server struct {
	Repo Repo
}

func (s *Server) PostUserLogin(ctx echo.Context) error {
	panic("implement me")
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
