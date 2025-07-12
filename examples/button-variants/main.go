// Package main demonstrates button variants with gio-shadcn.
package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"

	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/theme"
)

// layoutButtonRow creates a horizontal row of buttons with spacing.
func layoutButtonRow(gtx layout.Context, th *theme.Theme, buttons ...*button.Button) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return buttons[0].Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return buttons[1].Layout(gtx, th)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return buttons[2].Layout(gtx, th)
		}),
	)
}

func main() {
	go func() {
		var w app.Window
		w.Option(app.Title("Button Variants Example"))

		th := theme.New()

		// Create buttons with different variants
		buttons := []*button.Button{
			button.NewButton(
				button.WithText("Default"),
				button.WithVariant(theme.VariantDefault),
				button.WithOnClick(func() { println("Default clicked") }),
			),
			button.NewButton(
				button.WithText("Secondary"),
				button.WithVariant(theme.VariantSecondary),
				button.WithOnClick(func() { println("Secondary clicked") }),
			),
			button.NewButton(
				button.WithText("Destructive"),
				button.WithVariant(theme.VariantDestructive),
				button.WithOnClick(func() { println("Destructive clicked") }),
			),
			button.NewButton(
				button.WithText("Outline"),
				button.WithVariant(theme.VariantOutline),
				button.WithOnClick(func() { println("Outline clicked") }),
			),
			button.NewButton(
				button.WithText("Ghost"),
				button.WithVariant(theme.VariantGhost),
				button.WithOnClick(func() { println("Ghost clicked") }),
			),
			button.NewButton(
				button.WithText("Link"),
				button.WithVariant(theme.VariantLink),
				button.WithOnClick(func() { println("Link clicked") }),
			),
		}

		var ops op.Ops

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				return
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layoutButtonRow(gtx, th, buttons[0], buttons[1], buttons[2])
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layoutButtonRow(gtx, th, buttons[3], buttons[4], buttons[5])
						}),
					)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
