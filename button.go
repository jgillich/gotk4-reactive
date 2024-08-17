package reactive

import (
	"context"
	"errors"
	"fmt"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Button struct {
	model[*gtk.Button]
	Label   string
	Clicked func()
}

func (model *Button) CreateButton(ctx context.Context) (*gtk.Button, error) {
	w := gtk.NewButton()
	if err := model.Update(ctx, w); err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return w, nil
}

func (model *Button) Create(ctx context.Context) (gtk.Widgetter, error) {
	return model.CreateButton(ctx)
}

func (model *Button) Update(ctx context.Context, w gtk.Widgetter) error {
	button, ok := w.(*gtk.Button)
	if !ok {
		return errors.New("widget is not button")
	}

	button.SetLabel(model.Label)

	if h := glib.Bounded[glib.SignalHandle](button); h != nil {
		button.HandlerDisconnect(*h)
	}
	handle := button.ConnectClicked(model.Clicked)
	glib.Bind(button, handle)

	return nil
}
