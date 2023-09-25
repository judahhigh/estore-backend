package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	RegisterUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	RegisterUserResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	LoginUserResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RefreshUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	RefreshUserResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UnregisterUserRequest struct {
		Id string `json:"id"`
	}
	UnregisterUserResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type AuthorizationKey string

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decRegisterReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decLoginReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decRefreshReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decUnregisterReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UnregisterUserRequest
	vars := mux.Vars(r)
	req = UnregisterUserRequest{
		Id: vars["id"],
	}
	return req, nil
}
