// Package main demonstrates basic button usage with gio-shadcn.
package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"

	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

func main() {
	go func() {
		var w app.Window
		w.Option(app.Title("Basic Button Example"))

		th := theme.New()

		// Create button with functional options
		btn := button.NewButton(
			button.WithText("Click me"),
			button.WithVariant(theme.VariantDefault),
			button.WithOnClick(func() {
				println("Button clicked!")
			}),
		)

		var ops op.Ops

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				return
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return btn.Layout(gtx, th)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
