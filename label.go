package reactive

import (
	"context"
	"errors"
	"fmt"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Label struct {
	model[*gtk.Label]
	Text string
}

func (model *Label) CreateLabel(ctx context.Context) (*gtk.Label, error) {
	w := gtk.NewLabel(model.Text)
	if err := model.Update(ctx, w); err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return w, nil
}

func (model *Label) Create(ctx context.Context) (gtk.Widgetter, error) {
	return model.CreateLabel(ctx)
}

func (model *Label) Update(ctx context.Context, w gtk.Widgetter) error {
	label, ok := w.(*gtk.Label)
	if !ok {
		return errors.New("widget is not label")
	}

	label.SetText(model.Text)

	return nil
}
