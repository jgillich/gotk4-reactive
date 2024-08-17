package reactive

import (
	"context"
	"reflect"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Model interface {
	Type() reflect.Type
	Create(ctx context.Context) (gtk.Widgetter, error)
	Update(ctx context.Context, widget gtk.Widgetter) error
	Component() Component
}

type model[T gtk.Widgetter] struct{}

func (m *model[T]) Type() reflect.Type {
	return reflect.TypeFor[T]()
}

func (m *model[T]) Component() Component {
	return nil
}
