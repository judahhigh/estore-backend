package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	DeleteUserRequest struct {
		Id string `json:"id"`
	}
	DeleteUserResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UpdateUserRequest struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UpdateUserResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)
	req = GetUserRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeDeleteUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteUserRequest
	vars := mux.Vars(r)
	req = DeleteUserRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeUpdateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
