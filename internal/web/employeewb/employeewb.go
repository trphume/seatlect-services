package employeewb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/typedb"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetEmployeeBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PostEmployeeBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) DeleteEmployeeBusinessIdUsername(ctx echo.Context, businessId string, username string) error {
	panic("implement me")
}

type Repo interface {
	ListEmployee(ctx context.Context, businessId string) ([]typedb.Employee, error)
	CreateEmployee(ctx context.Context, businessId string, employee typedb.Employee) error
	DeleteEmployee(ctx context.Context, businessId string, username string) error
}
