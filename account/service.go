package account

import "context"

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (User, error)
	GetUser(ctx context.Context, id string) (User, error)
	DeleteUser(ctx context.Context, id string) (User, error)
}
