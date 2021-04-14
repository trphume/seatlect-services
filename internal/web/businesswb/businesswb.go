package businesswb

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/business_api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetBusiness(ctx echo.Context, params business_api.GetBusinessParams) error {
	business, err := s.Repo.SimpleListBusiness(ctx.Request().Context(), params.Status, params.Page)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := business_api.ListBusinessResponse{
		Businesses: typedbListToOapi(business),
	}

	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) GetBusinessBusinessId(ctx echo.Context, businessId string) error {
	business, err := s.Repo.GetBusinessById(ctx.Request().Context(), businessId)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Business not found with given id")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "ID is in an invalid format")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := typedbToOapi(*business)
	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) PatchBusinessBusinessId(ctx echo.Context, businessId string) error {
	var req business_api.UpdateBusinessRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	pId, err := primitive.ObjectIDFromHex(businessId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Bad ID")
	}

	request := typedb.Business{
		Id:           pId,
		BusinessName: *req.BusinessName,
		Type:         *req.Type,
		Tags:         *req.Tags,
		Description:  *req.Description,
		Location: typedb.Location{
			Type:        "Point",
			Coordinates: []float64{float64(*req.Location.Longitude), float64(*req.Location.Latitude)},
		},
		Address: *req.Address,
	}

	if err := s.Repo.UpdateBusinessById(ctx.Request().Context(), request); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Error fnding business with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	return ctx.String(http.StatusNoContent, "Business information updated successfully")
}

func (s *Server) PutBusinessBusinessIdDisplayImage(ctx echo.Context, businessId string) error {
	var req business_api.UpdateDisplayImageRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	img, err := s.Repo.UpdateBusinessDIById(ctx.Request().Context(), businessId, *req.DisplayImage)
	if err != nil {
		if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Bad argument")
		} else if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Can't find business id with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Internal error")
	}

	res := business_api.UpdateDisplayImageResponse{DisplayImage: createString(img)}
	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) PostBusinessBusinessIdImages(ctx echo.Context, businessId string) error {
	var req business_api.AppendImageRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	img, err := s.Repo.AppendBusinessImage(ctx.Request().Context(), businessId, *req.Image)
	if err != nil {
		if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Bad argument")
		} else if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Can't find business id with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Internal error")
	}

	res := business_api.AppendImageResponse{Image: createString(img)}
	return ctx.JSONPretty(http.StatusCreated, res, "  ")
}

func (s *Server) DeleteBusinessBusinessIdImagesPos(ctx echo.Context, businessId string, pos int) error {
	if err := s.Repo.RemoveBusinessImage(ctx.Request().Context(), businessId, pos); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Business not found with given id")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Invalid argument")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	return ctx.String(http.StatusNoContent, "Image deleted")
}

func (s *Server) GetBusinessBusinessIdMenu(ctx echo.Context, businessId string) error {
	menu, err := s.Repo.ListMenuItem(ctx.Request().Context(), businessId)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Business not found with given id")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "ID is in an invalid format")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := business_api.GetMenuResponse{Menu: typedbMenuToOapi(menu)}
	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) PostBusinessBusinessIdMenuitems(ctx echo.Context, businessId string) error {
	var req business_api.MenuItem
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	price, err := strconv.ParseFloat(*req.Price, 64)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Error with price format")
	}

	request := typedb.MenuItems{
		Name:        *req.Name,
		Description: *req.Description,
		Image:       *req.Image,
		Price:       price,
	}

	image, err := s.Repo.AppendMenuItem(ctx.Request().Context(), businessId, request)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Business not found with given id")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Invalid format")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := &business_api.AppendMenuItemResponse{Image: &image}
	return ctx.JSONPretty(http.StatusCreated, res, "  ")
}

func (s *Server) DeleteBusinessBusinessIdMenuitemsName(ctx echo.Context, businessId string, name string) error {
	if err := s.Repo.RemoveMenuItem(ctx.Request().Context(), businessId, name); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Business or menu item not found")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Argument format is missing or incorrect")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	return ctx.String(http.StatusNoContent, "Menu item deleted successfully")
}

func (s *Server) PatchBusinessBusinessIdStatus(ctx echo.Context, businessId string) error {
	var req business_api.UpdateStatusRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	if err := s.Repo.UpdateBusinessStatus(ctx.Request().Context(), businessId, *req.Status); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Business not found with given id")
		} else if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "ID is in an invalid format")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	return ctx.String(http.StatusNoContent, "Business status updated successfully")
}

type Repo interface {
	SimpleListBusiness(ctx context.Context, status int, page int) ([]typedb.Business, error)
	GetBusinessById(ctx context.Context, id string) (*typedb.Business, error)
	UpdateBusinessById(ctx context.Context, business typedb.Business) error
	UpdateBusinessDIById(ctx context.Context, id string, image string) (string, error)
	AppendBusinessImage(ctx context.Context, id string, image string) (string, error)
	RemoveBusinessImage(ctx context.Context, id string, pos int) error
	ListMenuItem(ctx context.Context, id string) ([]typedb.MenuItems, error)
	AppendMenuItem(ctx context.Context, id string, item typedb.MenuItems) (string, error)
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

func typedbMenuToOapi(menu []typedb.MenuItems) *[]business_api.MenuItem {
	res := make([]business_api.MenuItem, len(menu))
	for i, m := range menu {
		res[i] = typedbMenuItemsToOapi(m)
	}

	return &res
}

func typedbMenuItemsToOapi(m typedb.MenuItems) business_api.MenuItem {
	return business_api.MenuItem{
		Description: createString(m.Description),
		Image:       createString(m.Image),
		Name:        createString(m.Name),
		Price:       createString(fmt.Sprintf("%.2f", m.Price)),
	}
}

func createString(s string) *string {
	return &s
}

func createFloat32(f float32) *float32 {
	return &f
}
