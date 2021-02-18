package usermb

import (
	"context"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/genproto/commonpb"
	"github.com/tphume/seatlect-services/internal/genproto/userpb"
	"github.com/tphume/seatlect-services/internal/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Server struct {
	repo Repo

	userpb.UnimplementedUserServiceServer
}

func (s *Server) SignIn(ctx context.Context, req *userpb.SignInRequest) (*userpb.SignInResponse, error) {
	if !validation.ValidUsername(req.Username) || !validation.ValidPassword(req.Password) {
		return nil, status.Error(codes.InvalidArgument, "Argument is not valid")
	}

	customer := &typedb.Customer{Username: req.Username}

	token, err := s.repo.AuthenticateCustomer(ctx, customer, req.Password)
	if err != nil {
		if err == commonErr.INTERNAL {
			return nil, status.Error(codes.Internal, "Database error")
		}

		return nil, status.Error(codes.NotFound, "Username and password does not match")
	}

	return &userpb.SignInResponse{
		Token: token,
		User: &commonpb.User{
			Username: customer.Username,
			Dob:      customer.Dob.String(),
			Favorite: customer.Favorite},
	}, nil
}

func (s *Server) SignUp(ctx context.Context, req *userpb.SignUpRequest) (*userpb.SignUpResponse, error) {
	if !validation.ValidUsername(req.Username) || !validation.ValidPassword(req.Password) || !validation.ValidEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, "Argument is not valid")
	}

	// format request to correct type
	iso8601 := "2006-01-02T15:04:05-0700"
	dob, err := time.Parse(iso8601, req.Dob)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Argument is not valid")
	}

	customer := &typedb.Customer{Username: req.Username, Email: req.Email, Dob: dob}

	token, err := s.repo.CreateCustomer(ctx, customer, req.Password)
	if err != nil {
		if err == commonErr.INTERNAL {
			return nil, status.Error(codes.Internal, "Database error")
		}

		return nil, status.Error(codes.AlreadyExists, "User with that credentials already exist")
	}

	return &userpb.SignUpResponse{
		Token: token,
		User: &commonpb.User{
			Username: customer.Username,
			Dob:      customer.Dob.String(),
			Favorite: customer.Favorite},
	}, nil
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
	AuthenticateCustomer(ctx context.Context, customer *typedb.Customer, password string) (string, error)

	// customer is an out parameter - values will be overwritten
	// will return a user token
	CreateCustomer(ctx context.Context, customer *typedb.Customer, password string) (string, error)

	AppendFavorite(ctx context.Context, businessId string) error
	RemoveFavorite(ctx context.Context, businessId string) error
}
