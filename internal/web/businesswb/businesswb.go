package businesswb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/business_api"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetBusiness(ctx echo.Context, params business_api.GetBusinessParams) error {
	panic("implement me")
}

func (s *Server) GetBusinessBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PatchBusinessBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PutBusinessBusinessIdDisplayImage(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PostBusinessBusinessIdImages(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) DeleteBusinessBusinessIdImagesPos(ctx echo.Context, businessId string, pos int) error {
	panic("implement me")
}

func (s *Server) GetBusinessBusinessIdMenu(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) PostBusinessBusinessIdMenuitems(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) DeleteBusinessBusinessIdMenuitemsName(ctx echo.Context, businessId string, name string) error {
	panic("implement me")
}

func (s *Server) PatchBusinessBusinessIdStatus(ctx echo.Context, businessId string) error {
	panic("implement me")
}

type Repo interface {
	SimpleListBusiness(ctx context.Context, status int, page int) ([]typedb.Business, error)
	GetBusinessById(ctx context.Context, id string) (typedb.Business, error)
	UpdateBusinessById(ctx context.Context, business *typedb.Business) error
	UpdateBusinessDIById(ctx context.Context, id string, image string) error
	AppendBusinessImage(ctx context.Context, id string, image string) error
	RemoveBusinessImage(ctx context.Context, id string, pos int) error
	ListMenuItem(ctx context.Context, id string) ([]typedb.MenuItems, error)
	AppendMenuItem(ctx context.Context, id string, item typedb.MenuItems) error
	RemoveMenuItem(ctx context.Context, id string, name string) error
	UpdateBusinessStatus(ctx context.Context, id string, status int) error
}
