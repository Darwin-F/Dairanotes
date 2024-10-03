package entities

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type MethodsInterface interface {
	Store(ctx context.Context, notes Note) error
	Index(ctx context.Context, userID int64) ([]Note, error)
	Show(ctx context.Context, noteID int64) (*Note, error)
	Update(ctx context.Context, noteID int64, notes Note) error
	Destroy(ctx context.Context, noteID int64) error
}

type Methods struct {
	DB *sqlx.DB
}

func NewMethods(db *sqlx.DB) *Methods {
	return &Methods{
		DB: db,
	}
}

func (n *Methods) Store(ctx context.Context, notes Note) error {
	_, err := n.DB.ExecContext(ctx, "INSERT INTO notes (user_id, title, content) VALUES (?, ?, ?)", notes.UserID, notes.Title, notes.Content)
	if err != nil {
		return err
	}
	return nil
}

func (n *Methods) Index(ctx context.Context, userID int64) ([]Note, error) {
	var notes []Note

	rows, err := n.DB.QueryContext(ctx, "SELECT title, content FROM notes WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note Note

		err = rows.Scan(&note.Title, &note.Content)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, rows.Err()
}

func (n *Methods) Show(ctx context.Context, noteID int64) (*Note, error) {
	var note Note

	err := n.DB.QueryRowContext(ctx, "SELECT title,content  FROM notes WHERE id = ?", noteID).Scan(&note.Title, &note.Content)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (n *Methods) Update(ctx context.Context, noteID int64, notes Note) error {
	_, err := n.DB.ExecContext(ctx, "UPDATE notes SET title = ?, content = ? WHERE id = ?", notes.Title, notes.Content, noteID)
	if err != nil {
		return err
	}

	return nil
}

func (n *Methods) Destroy(ctx context.Context, noteID int64) error {
	_, err := n.DB.ExecContext(ctx, "DELETE FROM notes WHERE id = ?", noteID)
	if err != nil {
		return err
	}

	return nil
}
