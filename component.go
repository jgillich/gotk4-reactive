package reactive

import (
	"context"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/getseabird/seabird/internal/ctxt"
)

type Component interface {
	Update(ctx context.Context, message any, ch chan<- any) bool
	View(ctx context.Context, ch chan<- any) Model
}

type ComponentModel struct {
	model[*adw.Bin]
	component Component
}

func (c *ComponentModel) Create(ctx context.Context) (gtk.Widgetter, error) {
	node := ctxt.MustFrom[*Node](ctx)
	return c.component.View(ctx, node.ch).Create(ctx)
}

func (c *ComponentModel) Update(ctx context.Context, w gtk.Widgetter) error {
	node := ctxt.MustFrom[*Node](ctx)

	if component := glib.Bounded[Component](w); component != nil {
		c.component = *component
	} else {
		glib.Bind(w, c.Component)
	}

	return c.component.View(ctx, node.ch).Update(ctx, w)
}

func (m *ComponentModel) Component() Component {
	return m.component
}

func CreateComponent(component Component) Model {
	return &ComponentModel{component: component}
}
