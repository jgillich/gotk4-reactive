package reactive

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/getseabird/seabird/internal/ctxt"
)

type Box struct {
	model[*gtk.Box]
	Orientation gtk.Orientation
	Spacing     int
	Children    []Model
	Margin      [4]int
}

func (model *Box) CreateBox(ctx context.Context) (*gtk.Box, error) {
	w := gtk.NewBox(model.Orientation, model.Spacing)
	if err := model.Update(ctx, w); err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return w, nil
}

func (model *Box) Create(ctx context.Context) (gtk.Widgetter, error) {
	return model.CreateBox(ctx)
}

func (model *Box) Update(ctx context.Context, w gtk.Widgetter) (err error) {
	node := ctxt.MustFrom[*Node](ctx)
	box, ok := w.(*gtk.Box)
	if !ok {
		return errors.New("widget is not box")
	}

	next := box.FirstChild()
	for _, child := range model.Children {
		if next == nil {
			new := node.CreateChild(child)
			// new, err := child.Create(ctx)
			// if err != nil {
			// 	return err
			// }
			box.Append(new)
			continue
		}

		if child.Type() == reflect.TypeOf(next) {
			child.Update(ctx, next)
			next = gtk.BaseWidget(next).NextSibling()
		} else {
			new := node.CreateChild(child)
			// new, err := child.Create(ctx)
			// if err != nil {
			// 	return err
			// }
			box.InsertBefore(w, new)
			node.RemoveChild(w)
			box.Remove(w)
			next = gtk.BaseWidget(next).NextSibling()
		}
	}

	for {
		if next == nil {
			break
		}
		sibling := gtk.BaseWidget(next).NextSibling()
		box.Remove(next)
		next = sibling
	}

	box.SetSpacing(model.Spacing)
	box.SetMarginTop(model.Margin[0])
	box.SetMarginEnd(model.Margin[1])
	box.SetMarginBottom(model.Margin[2])
	box.SetMarginStart(model.Margin[3])

	return nil
}

func (model *Box) ID() string {
	return ""
}
