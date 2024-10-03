package entities

import "time"

type Note struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"-"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
