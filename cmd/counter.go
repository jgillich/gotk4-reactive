package main

import (
	"context"
	"fmt"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	r "github.com/jgillich/gotk4-reactive"
)

type Increment struct{}

type SampleComponent struct {
	counter int
}

func (c *SampleComponent) Update(ctx context.Context, message any, ch chan<- any) bool {
	switch message.(type) {
	case Increment:
		c.counter++
		return true
	default:
		return false
	}
}

func (c *SampleComponent) View(ctx context.Context, ch chan<- any) r.Model {
	return &r.Box{
		Orientation: gtk.OrientationVertical,
		Spacing:     5,
		Children: []r.Model{
			&r.Label{
				Text: fmt.Sprintf("Clicked %d times", c.counter),
			},
			&r.Button{
				Label: "Click me",
				Clicked: func() {
					ch <- Increment{}
				},
			},
		},
	}
}
