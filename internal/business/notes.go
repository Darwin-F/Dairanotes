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
	Store(ctx context.Context, note entities.Note) error
	Index(ctx context.Context, userID int64) ([]entities.Note, error)
	Show(ctx context.Context, noteID int64) (*entities.Note, error)
	Update(ctx context.Context, noteID int64, note entities.Note) error
	Destroy(ctx context.Context, noteID int64) error
}

func (n *NoteBusiness) Store(ctx context.Context, note entities.Note) error {
	return n.Methods.Store(ctx, note)
}

func (n *NoteBusiness) Index(ctx context.Context, userID int64) ([]entities.Note, error) {
	return n.Methods.Index(ctx, userID)
}

func (n *NoteBusiness) Show(ctx context.Context, noteID int64) (*entities.Note, error) {
	return n.Methods.Show(ctx, noteID)
}

func (n *NoteBusiness) Update(ctx context.Context, noteID int64, note entities.Note) error {
	return n.Methods.Update(ctx, noteID, note)
}

func (n *NoteBusiness) Destroy(ctx context.Context, noteID int64) error {
	return n.Methods.Destroy(ctx, noteID)
}
