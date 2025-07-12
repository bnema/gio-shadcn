// Package main demonstrates theme switching with gio-shadcn.
package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"

	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/theme"
)

func main() {
	go func() {
		var w app.Window
		w.Option(app.Title("Theme Switching Example"))

		th := theme.New()

		// Create theme toggle button
		themeToggle := button.NewButton(
			button.WithText("Toggle Dark Mode"),
			button.WithVariant(theme.VariantOutline),
			button.WithOnClick(func() {
				th.ToggleDark()
				if th.IsDark {
					println("Switched to dark mode")
				} else {
					println("Switched to light mode")
				}
			}),
		)

		// Create sample components to show theme changes
		primaryBtn := button.NewButton(
			button.WithText("Primary"),
			button.WithVariant(theme.VariantDefault),
			button.WithOnClick(func() { println("Primary clicked") }),
		)

		secondaryBtn := button.NewButton(
			button.WithText("Secondary"),
			button.WithVariant(theme.VariantSecondary),
			button.WithOnClick(func() { println("Secondary clicked") }),
		)

		destructiveBtn := button.NewButton(
			button.WithText("Destructive"),
			button.WithVariant(theme.VariantDestructive),
			button.WithOnClick(func() { println("Destructive clicked") }),
		)

		mainCard := card.NewCard(
			card.WithCardPadding(layout.Inset{Top: 32, Right: 32, Bottom: 32, Left: 32}),
		)

		var ops op.Ops

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				return
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return mainCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							// Theme toggle
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return themeToggle.Layout(gtx, th)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Spacer{Height: th.Spacing.Space8}.Layout(gtx)
							}),

							// Sample buttons to show theme changes
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return primaryBtn.Layout(gtx, th)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return secondaryBtn.Layout(gtx, th)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return destructiveBtn.Layout(gtx, th)
									}),
								)
							}),
						)
					})
				})

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
