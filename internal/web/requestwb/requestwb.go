package requestwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/request_api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Server struct {
	Repo    Repo
	BusRepo BusRepo
}

func (s *Server) GetRequest(ctx echo.Context, params request_api.GetRequestParams) error {
	requests, err := s.Repo.ListRequest(ctx.Request().Context(), params.Page)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := request_api.ListRequestResponse{
		Request: typedbListToOapi(requests),
	}

	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) GetRequestBusinessId(ctx echo.Context, businessId string) error {
	pId, err := primitive.ObjectIDFromHex(businessId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Bad ID")
	}

	request := &typedb.Request{Id: pId}
	if err := s.Repo.GetRequestById(ctx.Request().Context(), request); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Could not find change request with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := typedbToOapi(*request)
	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) PostRequestBusinessId(ctx echo.Context, businessId string) error {
	var req request_api.ChangeRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "Error binding request body")
	}

	pId, err := primitive.ObjectIDFromHex(businessId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Bad ID")
	}

	request := &typedb.Request{
		Id:           pId,
		BusinessName: *req.BusinessName,
		Type:         *req.Type,
		Tags:         *req.Tags,
		Description:  *req.Description,
		Location: typedb.Location{
			Type:        "Point",
			Coordinates: []float64{float64(*req.Location.Longitude), float64(*req.Location.Latitude)},
		},
		Address:   *req.Type,
		CreatedAt: time.Now(),
	}

	if err = s.Repo.CreateRequest(ctx.Request().Context(), request); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Error fnding business with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	return ctx.String(http.StatusCreated, "Business change request created successfully")
}

func (s *Server) DeleteRequestBusinessId(ctx echo.Context, businessId string) error {
	if err := s.Repo.DeleteRequest(ctx.Request().Context(), businessId); err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Error change request of business with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	return ctx.String(http.StatusNoContent, "Business change request rejected successfully")
}

func (s *Server) PostRequestBusinessIdApprove(ctx echo.Context, businessId string) error {
	_, err := s.Repo.ApproveRequest(ctx.Request().Context(), businessId)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Error change request of business with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	return ctx.String(http.StatusNoContent, "Business change request approved successfully")
}

type Repo interface {
	ListRequest(ctx context.Context, page int) ([]typedb.Request, error)
	ApproveRequest(ctX context.Context, id string) (string, error)
	GetRequestById(ctx context.Context, request *typedb.Request) error
	CreateRequest(ctx context.Context, request *typedb.Request) error
	DeleteRequest(ctx context.Context, id string) error
}

type BusRepo interface {
	GetBusinessById(ctx context.Context, id string, withMenu bool) (*typedb.Business, error)
}

// Helper functions
func typedbListToOapi(req []typedb.Request) *[]request_api.ChangeRequest {
	res := make([]request_api.ChangeRequest, len(req))
	for i, r := range req {
		res[i] = typedbToOapi(r)
	}

	return &res
}

func typedbToOapi(r typedb.Request) request_api.ChangeRequest {
	return request_api.ChangeRequest{
		Id:           createString(r.Id.Hex()),
		Address:      createString(r.Address),
		BusinessName: createString(r.BusinessName),
		Description:  createString(r.Description),
		Location: &request_api.Location{
			Longitude: createFloat32(float32(r.Location.Coordinates[0])),
			Latitude:  createFloat32(float32(r.Location.Coordinates[1])),
		},
		Tags: &r.Tags,
		Type: createString(r.Type),
	}
}

// Helper function
func createString(s string) *string {
	return &s
}

func createFloat32(f float32) *float32 {
	return &f
}
