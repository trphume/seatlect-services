package usermb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/userpb"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *Server) SignIn(context.Context, *userpb.SignInRequest) (*userpb.SignInResponse, error) {
	// TODO: Implement function
	panic("To be implemented")
}

func (s *Server) SignUp(context.Context, *userpb.SignUpRequest) (*userpb.SignUpResponse, error) {
	// TODO: Implement function
	panic("To be implemented")
}

func (s *Server) AddFavorite(context.Context, *userpb.AddFavoriteRequest) (*userpb.AddFavoriteResponse, error) {
	// TODO: Implement function
	panic("To be implemented")
}

func (s *Server) RemoveFavorite(context.Context, *userpb.RemoveFavoriteRequest) (*userpb.RemoveFavoriteResponse, error) {
	// TODO: Implement function
	panic("To be implemented")
}

type Repo interface {
	// customer is an out parameter - values will be overwritten
	// will return a user token
	AuthenticateCustomer(ctx *context.Context, customer *typedb.Customer, password string) (string, error)

	// customer is an out parameter - values will be overwritten
	// will return a user token
	CreateCustomer(ctx *context.Context, customer *typedb.Customer, password string) (string, error)

	AppendFavorite(ctx *context.Context, businessId string) error
	RemoveFavorite(ctx *context.Context, businessId string) error
}
