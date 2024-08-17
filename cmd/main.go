package main

import (
	"context"
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	r "github.com/jgillich/gotk4-reactive"
)

func main() {
	gtk.Init()

	app := gtk.NewApplication("dev.skynomads.Seabird", gio.ApplicationFlagsNone)

	window := r.ApplicationWindow{
		Title:       "Hello World",
		Height:      300,
		Width:       400,
		Application: app,
		Child:       r.CreateComponent(&SampleComponent{}),

		// &r.Box{
		// 	Orientation: gtk.OrientationHorizontal,
		// 	Margin:      [4]int{10, 10, 10, 10},
		// 	Spacing:     10,
		// 	Children: []r.Model{
		// 		&r.Label{Text: "foo"},
		// 		&r.Label{Text: "bar"},
		// 	},
		// },
	}

	app.ConnectActivate(func() {
		w := r.NewTree(context.Background(), &window).(*gtk.ApplicationWindow)
		w.Present()

		// window.Child.(*r.Box).Children = []r.Model{
		// 	&r.Label{Text: "bar"},
		// 	&r.Label{Text: "baz"},
		// 	// &r.LiveLabel{Label: r.Label{Text: "RX"}},
		// 	&r.Label{Text: "baz"},
		// 	&r.Label{Text: "baz"},
		// }
		// if err := window.Update(ctx, w); err != nil {
		// 	log.Fatal(err)
		// }
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}
