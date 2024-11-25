package business

import (
	"context"
	"dairanotes/internal/entities"
)

type UserBusiness struct {
	methods entities.UserMethodsInterface
}

func NewUserBusiness(userMethods entities.UserMethodsInterface) *UserBusiness {
	return &UserBusiness{methods: userMethods}
}

type UserBusinessInterface interface {
	Store(ctx context.Context, newUser entities.User) error
	Update(ctx context.Context, userID int64, user entities.User) error
	Destroy(ctx context.Context, userID int64) error
	GetUserID(ctx context.Context, username string) (int64, error)
}

func (u *UserBusiness) Store(ctx context.Context, newUser entities.User) error {
	return u.methods.Store(ctx, newUser)
}

func (u *UserBusiness) Update(ctx context.Context, userID int64, user entities.User) error {
	return u.methods.Update(ctx, userID, user)
}

func (u *UserBusiness) Destroy(ctx context.Context, userID int64) error {
	return u.methods.Destroy(ctx, userID)
}

func (u *UserBusiness) GetUserID(ctx context.Context, username string) (int64, error) {
	return u.methods.GetUserID(ctx, username)
}
