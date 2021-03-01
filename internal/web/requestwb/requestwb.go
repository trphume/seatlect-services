package requestwb

import (
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/gen_openapi/request_api"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetRequest(ctx echo.Context, params request_api.GetRequestParams) error {
	panic("implement me")
}

func (s *Server) GetRequestBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PostRequestBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PostRequestBusinessIdApprove(ctx echo.Context, businessId string) error {
	panic("implement me")
}

type Repo interface {
}
