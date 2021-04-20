package businessmb

import (
	"context"
	"fmt"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/businesspb"
	"github.com/tphume/seatlect-services/internal/genproto/commonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const iso8601 = "2006-01-02T15:04:05-0700"

type Server struct {
	Repo Repo

	businesspb.UnimplementedBusinessServiceServer
}

func (s *Server) ListBusiness(ctx context.Context, req *businesspb.ListBusinessRequest) (*businesspb.ListBusinessResponse, error) {
	searchParams := typedb.ListBusinessParams{
		Limit: req.Limit,
		Sort:  int32(req.Sort),
		Name:  req.Name,
		Type:  req.Type,
		Location: typedb.Location{
			Type:        "Point",
			Coordinates: []float64{req.Location.Longitude, req.Location.Latitude},
		},
	}

	// Call repo method
	businesses, err := s.Repo.ListBusiness(ctx, searchParams)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Convert typedb.Business to proto definition type
	res := typedbBusinessListToCommonpb(businesses)

	return &businesspb.ListBusinessResponse{Businesses: res}, nil
}

func (s *Server) ListBusinessById(ctx context.Context, req *businesspb.ListBusinessByIdRequest) (*businesspb.ListBusinessByIdResponse, error) {
	businesses, err := s.Repo.ListBusinessByIds(ctx, req.Ids)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Convert typedb.Business to proto definition type
	res := typedbBusinessListToCommonpb(businesses)

	return &businesspb.ListBusinessByIdResponse{Businesses: res}, nil
}

func (s *Server) GetBusinessById(ctx context.Context, req *businesspb.GetBusinessByIdRequest) (*businesspb.GetBusinessByIdResponse, error) {
	business, err := s.Repo.GetBusinessById(ctx, req.Id, true)
	if err != nil {
		if err == commonErr.NOTFOUND {
			return nil, status.Error(codes.NotFound, "Business not found with given id")
		}

		return nil, status.Error(codes.Internal, "Database error")
	}

	// Convert typedb.Business to proto definition type
	res := typedbBusinessToCommonpb(*business)

	return &businesspb.GetBusinessByIdResponse{Business: res}, nil
}

type Repo interface {
	ListBusiness(ctx context.Context, searchParams typedb.ListBusinessParams) ([]typedb.Business, error)
	ListBusinessByIds(ctx context.Context, ids []string) ([]typedb.Business, error)
	GetBusinessById(ctx context.Context, id string, withMenu bool) (*typedb.Business, error)
}

// Helper function
func typedbBusinessListToCommonpb(businesses []typedb.Business) []*commonpb.Business {
	res := make([]*commonpb.Business, len(businesses))
	for i, b := range businesses {
		res[i] = typedbBusinessToCommonpb(b)
	}

	return res
}

func typedbBusinessToCommonpb(b typedb.Business) *commonpb.Business {
	res := &commonpb.Business{
		XId:         b.Id.Hex(),
		Name:        b.BusinessName,
		Type:        b.Type,
		Tags:        b.Tags,
		Description: b.Description,
		Location: &commonpb.Latlng{
			Latitude:  b.Location.Coordinates[1],
			Longitude: b.Location.Coordinates[0],
		},
		Address:      b.Address,
		DisplayImage: b.DisplayImage,
		Images:       b.Images,
		Menu:         MenuItemsToProto(b.Menu),
	}

	return res
}

func MenuItemsToProto(mi []typedb.MenuItems) []*commonpb.MenuItem {
	res := make([]*commonpb.MenuItem, len(mi))
	for i, d := range mi {
		res[i] = &commonpb.MenuItem{
			Name:        d.Name,
			Description: d.Description,
			Image:       d.Image,
			Price:       fmt.Sprintf("%.2f", d.Price),
		}
	}

	return res
}
