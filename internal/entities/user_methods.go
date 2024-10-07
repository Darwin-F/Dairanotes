package entities

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type UserMethodsInterface interface {
	Store(ctx context.Context, newUser User) error
	Update(ctx context.Context, userID int64, user User) error
	Destroy(ctx context.Context, userID int64) error
}

type UserMethods struct {
	DB *sqlx.DB
}

func NewUserMethods(db *sqlx.DB) *UserMethods {
	return &UserMethods{
		DB: db,
	}
}

func (u *UserMethods) Store(ctx context.Context, newUser User) error {
	_, err := u.DB.ExecContext(ctx, "INSERT INTO users (name, email, password) VALUES (?, ?, ?)", newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserMethods) Update(ctx context.Context, userID int64, user User) error {
	_, err := u.DB.ExecContext(ctx, "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?", user.Username, user.Email, user.Password, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserMethods) Destroy(ctx context.Context, userID int64) error {
	_, err := u.DB.ExecContext(ctx, "DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}
