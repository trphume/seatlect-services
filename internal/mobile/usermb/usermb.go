package usermb

import (
	"context"
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
