package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	DeleteUser endpoint.Endpoint
	UpdateUser endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
		DeleteUser: makeDeleteUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		user, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, err
	}
}

func makeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		user, err := s.DeleteUser(ctx, req.Id)
		return DeleteUserResponse{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, err
	}
}

func makeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		user, err := s.UpdateUser(ctx, req.Id, req.Email, req.Password)
		return UpdateUserResponse{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, err
	}
}
