package requestwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/request_api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetRequest(ctx echo.Context, params request_api.GetRequestParams) error {
	requests := make([]typedb.Request, 0)
	max, err := s.Repo.ListRequest(ctx.Request().Context(), params.Page, requests)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := request_api.ListRequestResponse{
		MaxPage: &max,
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
	panic("implement me")
}

func (s *Server) PostRequestBusinessIdApprove(ctx echo.Context, businessId string) error {
	panic("implement me")
}

type Repo interface {
	ListRequest(ctx context.Context, page int, requests []typedb.Request) (int, error)
	ApproveRequest(ctX context.Context, id string) error
	GetRequestById(ctx context.Context, request *typedb.Request) error
	CreateRequest(ctx context.Context, request *typedb.Request) error
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
