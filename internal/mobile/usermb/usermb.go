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

	customer := &typedb.Customer{Username: req.Username, Password: req.Password}

	token, err := s.repo.AuthenticateCustomer(ctx, customer)
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

	customer := &typedb.Customer{Username: req.Username, Email: req.Email, Dob: dob, Password: req.Password, Favorite: make([]string, 0)}

	token, err := s.repo.CreateCustomer(ctx, customer)
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

func (s *Server) AddFavorite(ctx context.Context, req *userpb.AddFavoriteRequest) (*userpb.AddFavoriteResponse, error) {
	id := ctx.Value("id").(string)
	if len(id) <= 0 {
		return nil, status.Error(codes.Unauthenticated, "ID is not valid")
	}

	if err := s.repo.AppendFavorite(ctx, id, req.BusinessId); err != nil {
		if err == commonErr.NOTFOUND {
			return nil, status.Error(codes.NotFound, "Could not find business with that id")
		}

		return nil, status.Error(codes.Internal, "Database error")
	}

	return &userpb.AddFavoriteResponse{}, nil
}

func (s *Server) RemoveFavorite(ctx context.Context, req *userpb.RemoveFavoriteRequest) (*userpb.RemoveFavoriteResponse, error) {
	id := ctx.Value("id").(string)
	if len(id) <= 0 {
		return nil, status.Error(codes.Unauthenticated, "ID is not valid")
	}

	if err := s.repo.RemoveFavorite(ctx, id, req.BusinessId); err != nil {
		if err == commonErr.NOTFOUND {
			return nil, status.Error(codes.NotFound, "Could not find business with that id")
		}

		return nil, status.Error(codes.Internal, "Database error")
	}

	return &userpb.RemoveFavoriteResponse{}, nil
}

type Repo interface {
	// customer is an out parameter - values will be overwritten
	// will return a user token
	AuthenticateCustomer(ctx context.Context, customer *typedb.Customer) (string, error)

	// customer is an out parameter - values will be overwritten
	// will return a user token
	CreateCustomer(ctx context.Context, customer *typedb.Customer) (string, error)

	AppendFavorite(ctx context.Context, customerId string, businessId string) error
	RemoveFavorite(ctx context.Context, customerId string, businessId string) error
}
