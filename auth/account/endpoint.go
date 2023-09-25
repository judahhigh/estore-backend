package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	RegisterUser   endpoint.Endpoint
	LoginUser      endpoint.Endpoint
	RefreshUser    endpoint.Endpoint
	UnregisterUser endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		RegisterUser:   makeRegisterUserEndpoint(s),
		LoginUser:      makeLoginUserEndpoint(s),
		RefreshUser:    makeRefreshUserEndpoint(s),
		UnregisterUser: makeUnregisterUserEndpoint(s),
	}
}

func makeRegisterUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterUserRequest)
		user, err := s.Register(ctx, req.Email, req.Password)
		return RegisterUserResponse{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, err
	}
}

func makeLoginUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginUserRequest)
		user, err := s.Login(ctx, req.Email, req.Password)
		return LoginUserResponse{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, err
	}
}

func makeRefreshUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RefreshUserRequest)
		user, err := s.Refresh(ctx, req.Email, req.Password)
		return RefreshUserResponse{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, err
	}
}

func makeUnregisterUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UnregisterUserRequest)
		user, err := s.Unregister(ctx, req.Id)
		return UnregisterUserResponse{
			Id:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, err
	}
}
