// Package main demonstrates form components with gio-shadcn.
package main

import (
	"image"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"

	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/theme"
)

func main() {
	go func() {
		var w app.Window
		w.Option(app.Title("Form Example"))

		th := theme.New()

		// Create form components
		nameInput := input.New()
		nameInput.Placeholder = "Enter your name"
		nameInput.Type = input.InputText
		nameInput.Label = "Name"
		nameInput.Required = true

		emailInput := input.New()
		emailInput.Placeholder = "Enter your email"
		emailInput.Type = input.InputEmail
		emailInput.Label = "Email"
		emailInput.Required = true

		submitBtn := button.NewButton(
			button.WithText("Submit"),
			button.WithVariant(theme.VariantDefault),
			button.WithOnClick(func() {
				println("Form submitted!")
				println("Name:", nameInput.Text())
				println("Email:", emailInput.Text())
			}),
		)

		resetBtn := button.NewButton(
			button.WithText("Reset"),
			button.WithVariant(theme.VariantSecondary),
			button.WithOnClick(func() {
				nameInput.SetText("")
				emailInput.SetText("")
				println("Form reset!")
			}),
		)

		formCard := card.NewCard(
			card.WithCardPadding(layout.Inset{Top: 24, Right: 24, Bottom: 24, Left: 24}),
		)

		var ops op.Ops

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				return
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return formCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							// Title
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{Bottom: th.Spacing.Space6}.Layout(gtx, func(_ layout.Context) layout.Dimensions {
									// You would use a Label component here
									return layout.Dimensions{Size: image.Point{X: 200, Y: 40}}
								})
							}),

							// Name input
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return nameInput.Layout(gtx, th)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
							}),

							// Email input
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return emailInput.Layout(gtx, th)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Spacer{Height: th.Spacing.Space6}.Layout(gtx)
							}),

							// Buttons
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return submitBtn.Layout(gtx, th)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return resetBtn.Layout(gtx, th)
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
