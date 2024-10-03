package business

import (
	"Dairanotes/internal/entities"
	"context"
)

type NoteBusiness struct {
	noteMethods entities.NoteMethodsInterface
}

func NewNoteBusiness(noteMethods entities.NoteMethodsInterface) *NoteBusiness {
	return &NoteBusiness{noteMethods: noteMethods}
}

type NoteBusinessInterface interface {
	CreateNote(ctx context.Context, note entities.Note) error
	GetNotes(ctx context.Context, userID int64) ([]entities.Note, error)
	GetNoteByID(ctx context.Context, noteID int64) (*entities.Note, error)
	UpdateNoteByID(ctx context.Context, noteID int64) error
	DeleteNoteByID(ctx context.Context, noteID int64) error
}

func (n *NoteBusiness) CreateNote(ctx context.Context, note entities.Note) error {
	return n.noteMethods.CreateNote(ctx, note)
}

func (n *NoteBusiness) GetNotes(ctx context.Context, userID int64) ([]entities.Note, error) {
	return n.noteMethods.GetNotes(ctx, userID)
}

func (n *NoteBusiness) GetNoteByID(ctx context.Context, noteID int64) (*entities.Note, error) {
	return n.noteMethods.GetNoteByID(ctx, noteID)
}

func (n *NoteBusiness) UpdateNoteByID(ctx context.Context, noteID int64, note entities.Note) error {
	return n.noteMethods.UpdateNoteByID(ctx, noteID, note)
}

func (n *NoteBusiness) DeleteNoteByID(ctx context.Context, noteID int64) error {
	return n.noteMethods.DeleteNoteByID(ctx, noteID)
}
