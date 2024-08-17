package reactive

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/getseabird/seabird/internal/ctxt"
)

type ApplicationWindow struct {
	model[*gtk.ApplicationWindow]
	Application *gtk.Application
	Title       string
	Child       Model
	Height      int
	Width       int
}

func (model *ApplicationWindow) CreateApplicationWindow(ctx context.Context) (*gtk.ApplicationWindow, error) {
	w := gtk.NewApplicationWindow(model.Application)
	if err := model.Update(ctx, w); err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return w, nil
}

func (model *ApplicationWindow) Create(ctx context.Context) (gtk.Widgetter, error) {
	return model.CreateApplicationWindow(ctx)
}

func (model *ApplicationWindow) Update(ctx context.Context, w gtk.Widgetter) error {
	node := ctxt.MustFrom[*Node](ctx)
	window, ok := w.(*gtk.ApplicationWindow)
	if !ok {
		return errors.New("widget is not application window")
	}

	if child := window.Child(); model.Child.Type() == reflect.TypeOf(child) {
		return model.Child.Update(ctx, child)
	} else {
		c := node.CreateChild(model.Child)
		// c, err := model.Child.Create(ctx)
		// if err != nil {
		// 	return err
		// }
		window.SetChild(c)
	}

	window.SetTitle(model.Title)
	window.SetDefaultSize(model.Width, model.Height)

	return nil
}
