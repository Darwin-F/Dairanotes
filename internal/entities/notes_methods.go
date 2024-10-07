package entities

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type NotesMethodsInterface interface {
	Store(ctx context.Context, notes Note) error
	Index(ctx context.Context, userID int64) ([]Note, error)
	Show(ctx context.Context, noteID int64) (*Note, error)
	Update(ctx context.Context, noteID int64, notes Note) error
	Destroy(ctx context.Context, noteID int64) error
}

type NotesMethods struct {
	DB *sqlx.DB
}

func NewNotesMethods(db *sqlx.DB) *NotesMethods {
	return &NotesMethods{
		DB: db,
	}
}

func (n *NotesMethods) Store(ctx context.Context, notes Note) error {
	_, err := n.DB.ExecContext(ctx, "INSERT INTO notes (user_id, title, content) VALUES (?, ?, ?)", notes.UserID, notes.Title, notes.Content)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotesMethods) Index(ctx context.Context, userID int64) ([]Note, error) {
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

func (n *NotesMethods) Show(ctx context.Context, noteID int64) (*Note, error) {
	var note Note

	err := n.DB.QueryRowContext(ctx, "SELECT title,content  FROM notes WHERE id = ?", noteID).Scan(&note.Title, &note.Content)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (n *NotesMethods) Update(ctx context.Context, noteID int64, notes Note) error {
	_, err := n.DB.ExecContext(ctx, "UPDATE notes SET title = ?, content = ? WHERE id = ?", notes.Title, notes.Content, noteID)
	if err != nil {
		return err
	}

	return nil
}

func (n *NotesMethods) Destroy(ctx context.Context, noteID int64) error {
	_, err := n.DB.ExecContext(ctx, "DELETE FROM notes WHERE id = ?", noteID)
	if err != nil {
		return err
	}

	return nil
}
