package adminwb

import "github.com/labstack/echo/v4"

type Server struct {
	Repo Repo
}

func (s *Server) PostAdminLogin(ctx echo.Context) error {
	panic("implement me")
}

type Repo interface {
}
