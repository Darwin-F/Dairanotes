package business

import (
	"context"
	"dairanotes/internal/entities"
)

type NoteBusiness struct {
	Methods entities.MethodsInterface
}

func NewNoteBusiness(noteMethods entities.MethodsInterface) *NoteBusiness {
	return &NoteBusiness{Methods: noteMethods}
}

type NoteBusinessInterface interface {
	CreateNote(ctx context.Context, note entities.Note) error
	GetNotes(ctx context.Context, userID int64) ([]entities.Note, error)
	GetNoteByID(ctx context.Context, noteID int64) (*entities.Note, error)
	UpdateNoteByID(ctx context.Context, noteID int64, note entities.Note) error
	DeleteNoteByID(ctx context.Context, noteID int64) error
}

func (n *NoteBusiness) CreateNote(ctx context.Context, note entities.Note) error {
	return n.Methods.CreateNote(ctx, note)
}

func (n *NoteBusiness) GetNotes(ctx context.Context, userID int64) ([]entities.Note, error) {
	return n.Methods.GetNotes(ctx, userID)
}

func (n *NoteBusiness) GetNoteByID(ctx context.Context, noteID int64) (*entities.Note, error) {
	return n.Methods.GetNoteByID(ctx, noteID)
}

func (n *NoteBusiness) UpdateNoteByID(ctx context.Context, noteID int64, note entities.Note) error {
	return n.Methods.UpdateNoteByID(ctx, noteID, note)
}

func (n *NoteBusiness) DeleteNoteByID(ctx context.Context, noteID int64) error {
	return n.Methods.DeleteNoteByID(ctx, noteID)
}
