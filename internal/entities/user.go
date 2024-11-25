package entities

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username" binding:"required"`
	Password   string    `json:"-" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	VerifiedAt time.Time `json:""`
	CreatedAt  time.Time `json:""`
	UpdatedAt  time.Time `json:""`
}
