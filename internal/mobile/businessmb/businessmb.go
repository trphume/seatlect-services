package businessmb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/businesspb"
	"github.com/tphume/seatlect-services/internal/genproto/commonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	repo Repo

	businesspb.UnimplementedBusinessServiceServer
}

func (s *Server) ListBusiness(ctx context.Context, request *businesspb.ListBusinessRequest) (*businesspb.ListBusinessResponse, error) {
	panic("implement me")
}

func (s *Server) ListBusinessById(ctx context.Context, req *businesspb.ListBusinessByIdRequest) (*businesspb.ListBusinessByIdResponse, error) {
	businesses, err := s.repo.ListBusinessByIds(ctx, req.Ids)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Convert typedb.Business to proto definition type
	res := make([]*commonpb.Business, len(businesses))
	for i, b := range businesses {
		res[i] = &commonpb.Business{
			XId:         b.Id.Hex(),
			Name:        b.BusinessName,
			Type:        b.Type,
			Tags:        b.Tags,
			Description: b.Description,
			Location: &commonpb.Latlng{
				Latitude:  b.Location.Coordinates[0],
				Longitude: b.Location.Coordinates[1],
			},
			Address:      b.Address,
			DisplayImage: b.DisplayImage,
			Images:       b.Images,
			Menu:         MenuItemsToProto(b.Menu),
		}
	}

	return &businesspb.ListBusinessByIdResponse{Businesses: res}, nil
}

type Repo interface {
	ListBusiness(ctx context.Context, searchParams typedb.ListBusinessParams) ([]typedb.Business, error)
	ListBusinessByIds(ctx context.Context, ids []string) ([]typedb.Business, error)
}

// Helper function
func MenuItemsToProto(mi []typedb.MenuItems) []*commonpb.MenuItem {
	res := make([]*commonpb.MenuItem, len(mi))
	for i, d := range mi {
		res[i] = &commonpb.MenuItem{
			Name:        d.Name,
			Description: d.Description,
			Image:       d.Image,
			Price:       d.Price.String(),
		}
	}

	return res
}
