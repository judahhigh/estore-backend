package account

import "context"

type User struct {
	ID       string `json:"id,omitempty" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUser(ctx context.Context, id string) (User, error)
	DeleteUser(ctx context.Context, id string) (User, error)
}
