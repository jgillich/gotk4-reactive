package reactive

import (
	"context"
	"slices"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/getseabird/seabird/internal/ctxt"
)

func NewTree(ctx context.Context, model Model) gtk.Widgetter {
	root := &Node{
		ch:  make(chan any),
		ctx: ctx,
	}
	ctx = ctxt.With[*Node](ctx, root)

	root.widget, _ = model.Create(ctx)

	go func() {
		for {
			root.message(<-root.ch, true)
		}
	}()

	return root.widget
}

type Node struct {
	ch     chan any
	parent *Node
	ctx    context.Context
	cancel context.CancelFunc
	// model     Model
	widget    gtk.Widgetter
	component Component
	children  []*Node
}

func (r *Node) Render(model Model) gtk.Widgetter {
	return nil
}

func (n *Node) CreateChild(model Model) gtk.Widgetter {
	node := &Node{parent: n, ch: n.ch}
	node.ctx, node.cancel = context.WithCancel(ctxt.With[*Node](n.ctx, node))

	node.widget, _ = model.Create(node.ctx)
	glib.Bind[*Node](node.widget, node)
	n.children = append(n.children, node)

	if c := model.Component(); c != nil {
		node.component = c
	}

	return node.widget
}

func (n *Node) RemoveChild(widget gtk.Widgetter) {
	node := *glib.Bounded[*Node](widget)
	node.cancel()
	if p := node.parent; p != nil {
		for i, n := range node.children {
			if n.widget == p.widget {
				slices.Delete(node.children, i, i+1)
				break
			}
		}
	}
}

func (n *Node) message(msg any, rerender bool) {
	if n.component != nil {
		if n.component.Update(n.ctx, msg, n.ch) && rerender {
			rerender = false
			defer n.component.View(n.ctx, n.ch).Update(n.ctx, n.widget)
		}
	}
	for _, c := range n.children {
		c.message(msg, rerender)
	}
}
