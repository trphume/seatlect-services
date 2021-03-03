package requestwb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/request_api"
	"net/http"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetRequest(ctx echo.Context, params request_api.GetRequestParams) error {
	requests := make([]typedb.Request, 0)
	max, err := s.Repo.ListRequest(ctx.Request().Context(), params.Page, requests)
	if err != nil {
		return ctx.JSONPretty(http.StatusInternalServerError, "Database error", "  ")
	}

	res := request_api.ListRequestResponse{
		MaxPage: &max,
		Request: typedbToOapi(requests),
	}

	return ctx.JSONPretty(http.StatusOK, res, "  ")
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
	ListRequest(ctx context.Context, page int, requests []typedb.Request) (int, error)
	ApproveRequest(ctX context.Context, id string) error
	GetRequestById(ctx context.Context, request *typedb.Request) error
	CreateRequest(ctx context.Context, request *typedb.Request) error
}

// Helper functions
func typedbToOapi(req []typedb.Request) *[]request_api.ChangeRequest {
	res := make([]request_api.ChangeRequest, len(req))
	for i, r := range req {
		res[i] = request_api.ChangeRequest{
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

	return &res
}

// Helper function
func createString(s string) *string {
	return &s
}

func createFloat32(f float32) *float32 {
	return &f
}
