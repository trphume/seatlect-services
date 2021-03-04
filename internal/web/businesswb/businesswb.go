package businesswb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/business_api"
	"net/http"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetBusiness(ctx echo.Context, params business_api.GetBusinessParams) error {
	business := make([]typedb.Business, 0)
	max, err := s.Repo.SimpleListBusiness(ctx.Request().Context(), params.Status, params.Page, business)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := business_api.ListBusinessResponse{
		Businesses: typedbListToOapi(business),
		MaxPage:    &max,
	}

	return ctx.JSONPretty(http.StatusOK, res, "  ")
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
	SimpleListBusiness(ctx context.Context, status int, page int, business []typedb.Business) (int, error)
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

// Helper function
func typedbListToOapi(business []typedb.Business) *[]business_api.Business {
	res := make([]business_api.Business, len(business))
	for i, b := range business {
		res[i] = typedbToOapi(b)
	}

	return &res
}

func typedbToOapi(b typedb.Business) business_api.Business {
	return business_api.Business{
		Id:           createString(b.Id.Hex()),
		Address:      createString(b.Address),
		BusinessName: createString(b.BusinessName),
		Description:  createString(b.Description),
		DisplayImage: createString(b.DisplayImage),
		Images:       &b.Images,
		Location: &business_api.Location{
			Longitude: createFloat32(float32(b.Location.Coordinates[0])),
			Latitude:  createFloat32(float32(b.Location.Coordinates[1])),
		},
		Tags: &b.Tags,
		Type: createString(b.Type),
	}
}

func createString(s string) *string {
	return &s
}

func createFloat32(f float32) *float32 {
	return &f
}
