package account

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (User, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	user, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}

	logger.Log("create user", id)

	return user, nil
}

func (s service) GetUser(ctx context.Context, id string) (User, error) {
	logger := log.With(s.logger, "method", "GetUser")

	user, err := s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}

	logger.Log("Get user", id)

	return user, nil
}

func (s service) DeleteUser(ctx context.Context, id string) (User, error) {
	logger := log.With(s.logger, "method", "DeleteUser")

	user, err := s.repository.DeleteUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}

	logger.Log("delete user", id)

	return user, nil
}
