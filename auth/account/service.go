package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"context"

	"net/http"
)

type Service interface {
	Register(ctx context.Context, email string, password string) (User, error)
	Login(ctx context.Context, email string, password string) (User, error)
	Refresh(ctx context.Context, email string, password string) (User, error)
	Unregister(ctx context.Context, id string) (User, error)
}

type accountService struct {
	logger log.Logger
}

func NewService(logger log.Logger) Service {
	return &accountService{
		logger: logger,
	}
}

func (s accountService) Register(ctx context.Context, email string, password string) (User, error) {
	logger := log.With(s.logger, "method", "Register")

	// User to be filled with response
	var user User

	// Get the authorization header from the request and pass through to the api
	r := ctx.Value(ctxRequestKey{}).(*http.Request)
	auth := r.Header.Get("Authorization")

	// Load the account api server details
	var server_details AccountApiServerDetails
	err := server_details.Load()
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}

	// Make a post to the server to register a new user, effectively creating a user
	post_body, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	response_body := bytes.NewBuffer(post_body)
	conn_string := fmt.Sprintf("%s://%s:%s/user", server_details.scheme, server_details.host, server_details.port)
	req, err := http.NewRequest("POST", conn_string, response_body)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)
	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}
	defer resp.Body.Close()
	// resp, err := http.Post(conn_string, "application/json", response_body)
	// if err != nil {
	// 	level.Error(logger).Log("err", err)
	// 	return user, err
	// }
	// defer resp.Body.Close()

	// Deserialize the body of the response into the user to return to the client
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}
	json.Unmarshal(body, &user)

	// Return the registered user information in the response
	logger.Log("Registered user", email)
	return user, nil
}

func (s accountService) Login(ctx context.Context, email string, password string) (User, error) {
	logger := log.With(s.logger, "method", "Register")
	logger.Log("Logged in user", email)
	return User{}, nil
}

func (s accountService) Refresh(ctx context.Context, email string, password string) (User, error) {
	logger := log.With(s.logger, "method", "Register")
	logger.Log("Refreshed token for user", email)
	return User{}, nil
}

func (s accountService) Unregister(ctx context.Context, id string) (User, error) {
	logger := log.With(s.logger, "method", "Unregister")

	// User to be filled with response
	var user User

	// Get the authorization header from the request and pass through to the api
	r := ctx.Value(ctxRequestKey{}).(*http.Request)
	auth := r.Header.Get("Authorization")

	// Load the account api server details
	var server_details AccountApiServerDetails
	err := server_details.Load()
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}

	// Make a delete request to the server to delete the user, effectively unregistering them
	client := &http.Client{}
	conn_string := fmt.Sprintf("%s://%s:%s/user/%s", server_details.scheme, server_details.host, server_details.port, id)
	println("\n", conn_string, "\n")
	req, err := http.NewRequest("DELETE", conn_string, nil)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}
	req.Header.Add("Authorization", auth)
	// send the request
	resp, err := client.Do(req)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}
	defer resp.Body.Close()

	// Deserialize the body of the response into the user to return to the client
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}
	json.Unmarshal(body, &user)

	// Return the unregistered user information in the response
	logger.Log("Unregistered user", user.Email)
	return user, nil
}

type AccountApiServerDetails struct {
	scheme string
	host   string
	port   string
}

type errorAccountApiServerDetails struct {
	s string
}

func (e *errorAccountApiServerDetails) Error() string {
	return e.s
}

func (serverDetails *AccountApiServerDetails) Load() error {
	host, result := os.LookupEnv("ACCOUNT_API_HOST")
	if !result {
		return &errorAccountApiServerDetails{"Cannot load account api server details from the environment."}
	}
	serverDetails.host = host

	scheme, result := os.LookupEnv("ACCOUNT_API_SCHEME")
	if !result {
		return &errorAccountApiServerDetails{"Cannot load account api server details from the environment."}
	}
	serverDetails.scheme = scheme

	port, result := os.LookupEnv("ACCOUNT_API_PORT")
	if !result {
		return &errorAccountApiServerDetails{"Cannot load account api server details from the environment."}
	}
	serverDetails.port = port

	return nil
}
