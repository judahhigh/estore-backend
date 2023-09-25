package account

import (
	"context"
	"errors"

	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
)

var ErrRepo = errors.New("unable to handle Repo Request")

type repo struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewRepo(db *sqlx.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) (User, error) {
	sql := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)`

	if user.Email == "" || user.Password == "" {
		return user, ErrRepo
	}

	println("\nUSER: ", user.ID, user.Email, user.Password, "\n")
	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (User, error) {
	user := User{}
	rows, err := repo.db.Queryx("SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return user, ErrRepo
	}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return user, ErrRepo
		}
	}

	return user, nil
}

func (repo *repo) GetUsers(ctx context.Context) ([]User, error) {
	users := []User{}
	rows, err := repo.db.Queryx("SELECT * FROM users")
	if err != nil {
		return users, ErrRepo
	}
	for rows.Next() {
		user := User{}
		err := rows.StructScan(&user)
		if err != nil {
			return users, ErrRepo
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *repo) DeleteUser(ctx context.Context, id string) (User, error) {
	user := User{}
	user.ID = id
	rows, read_err := repo.db.Queryx("SELECT * FROM users WHERE id=$1", id)
	if read_err != nil {
		return user, ErrRepo
	}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return user, ErrRepo
		}
	}

	sql := `DELETE FROM users WHERE id=$1`
	_, write_err := repo.db.ExecContext(ctx, sql, id)
	if write_err != nil {
		return user, write_err
	}
	return user, nil
}

func (repo *repo) UpdateUser(ctx context.Context, user User) (User, error) {
	// Check that actual values are being used to update data
	if user.Email == "" || user.Password == "" {
		return user, ErrRepo
	}

	// Attempt to perform the update.
	sql := `
		UPDATE users SET email=$1, password=$2 WHERE id=$3`

	println("\nUSER: ", user.ID, user.Email, user.Password, "\n")
	_, err := repo.db.ExecContext(ctx, sql, user.Email, user.Password, user.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}
