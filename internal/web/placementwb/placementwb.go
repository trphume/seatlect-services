package placementwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/typedb"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetPlacementBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PutPlacementBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

type Repo interface {
	GetPlacement(ctx context.Context, id string) (typedb.Placement, error)
	UpdatePlacement(ctx context.Context, id string, placement typedb.Placement) error
}
