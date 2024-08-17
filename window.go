package reactive

import (
	"context"
	"errors"
	"fmt"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Window struct {
	model[*gtk.Window]
	Title string
	Child Model
}

func (model *Window) CreateWindow(ctx context.Context) (*gtk.Window, error) {
	w := gtk.NewWindow()
	if err := model.Update(ctx, w); err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return w, nil
}

func (model *Window) Create(ctx context.Context) (gtk.Widgetter, error) {
	w, err := model.CreateWindow(ctx)
	return w, err
}

func (model *Window) Update(ctx context.Context, w gtk.Widgetter) error {
	window, ok := w.(*gtk.Window)
	if !ok {
		return errors.New("widget is not window")
	}

	// if child := window.Child(); model.Child.Is(child) {
	// 	return model.Child.Update(ctx, child)
	// } else {
	// 	c, err := model.Child.Create(ctx)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	window.SetChild(c)
	// }

	window.SetTitle(model.Title)

	return nil
}
