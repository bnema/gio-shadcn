/*
Package main provides a comprehensive demo application showcasing gio-shadcn components.

This demo application demonstrates all available gio-shadcn UI components in action,
including theming, interactivity, and responsive design. It serves as both a
functional example and a testing ground for the component library.

# Features

• Complete showcase of all available components
• Light/dark theme switching with live preview
• Interactive zoom functionality (Ctrl+/-, Ctrl+0)
• Custom frameless window with titlebar
• Responsive layout design
• Component interaction examples
• Theme integration demonstration

# Usage

Run the demo application:

	go run ./cmd/demo-app

# Components Demonstrated

• Button - All variants (default, destructive, outline, secondary, ghost, link)
• Card - Container with structured content
• Input - Text input with placeholder and validation
• Label - Typography elements (H1-H4, body text, small text)
• Titlebar - Custom window controls and branding

# Interactive Features

Keyboard shortcuts:
• Ctrl + Plus/Equals - Zoom in
• Ctrl + Minus - Zoom out
• Ctrl + 0 - Reset zoom to 100%

UI controls:
• Theme toggle button (light/dark mode switching)
• Zoom control buttons with current zoom display
• Window controls (minimize, maximize, close)
• Interactive component examples

# Architecture

The demo follows a clean architecture pattern:
• Theme management with runtime switching
• Component state management
• Event handling for keyboard and mouse
• Responsive layout with zoom support
• Window management for frameless design

This demo serves as the primary example for integrating gio-shadcn
components into real applications and demonstrates best practices
for theme usage, component composition, and user interaction.
*/
package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/card"
	"github.com/bnema/gio-shadcn/components/input"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/components/titlebar"
	"github.com/bnema/gio-shadcn/theme"
)

func main() {
	go func() {
		w := &app.Window{}
		w.Option(app.Title("Gio-shadcn Demo"))
		w.Option(app.Size(800, 600))
		w.Option(app.Decorated(false)) // Disable system title bar

		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

// updateWindowColors updates the window colors based on the current theme.
func updateWindowColors(w *app.Window, th *theme.Theme) {
	w.Option(app.NavigationColor(th.Colors.Background))
	w.Option(app.StatusColor(th.Colors.Background))
}

// layoutButtonRow renders a horizontal row of buttons with spacing.
func layoutButtonRow(gtx layout.Context, th *theme.Theme, buttons ...*button.Button) layout.Dimensions {
	children := make([]layout.FlexChild, 0, len(buttons)*2)
	for i, btn := range buttons {
		btn := btn // capture loop variable
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return btn.Layout(gtx, th)
		}))
		// Add spacing between buttons (except after the last one)
		if i < len(buttons)-1 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
			}))
		}
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
}

func run(w *app.Window) error {
	// Initialize theme
	th := theme.New()

	// Set initial window colors to match theme
	updateWindowColors(w, th)

	// Initialize title bar with secondary variant
	tb := titlebar.NewTitleBar(
		titlebar.WithTitle("Gio-shadcn Demo"),
		titlebar.WithWindow(w),
		titlebar.WithVariant(theme.VariantSecondary),
	)

	// Zoom state
	var zoomScale float32 = 1.0
	const minZoom = 0.5
	const maxZoom = 3.0
	const zoomStep = 0.1

	// Create zoom label first (needed by zoom buttons)
	zoomLabel := label.NewTypography("Zoom: 100%", label.Small, "")

	// Initialize components
	primaryBtn := button.New(button.Config{
		Text:    "Primary Button",
		Variant: theme.VariantDefault,
		Size:    theme.SizeDefault,
		OnClick: func() {
			log.Println("Primary button clicked!")
		},
	})

	destructiveBtn := button.New(button.Config{
		Text:    "Destructive",
		Variant: theme.VariantDestructive,
		Size:    theme.SizeDefault,
		OnClick: func() {
			log.Println("Destructive button clicked!")
		},
	})

	outlineBtn := button.New(button.Config{
		Text:    "Outline",
		Variant: theme.VariantOutline,
		Size:    theme.SizeDefault,
		OnClick: func() {
			log.Println("Outline button clicked!")
		},
	})

	secondaryBtn := button.New(button.Config{
		Text:    "Secondary",
		Variant: theme.VariantSecondary,
		Size:    theme.SizeDefault,
		OnClick: func() {
			log.Println("Secondary button clicked!")
		},
	})

	ghostBtn := button.New(button.Config{
		Text:    "Ghost",
		Variant: theme.VariantGhost,
		Size:    theme.SizeDefault,
		OnClick: func() {
			log.Println("Ghost button clicked!")
		},
	})

	linkBtn := button.New(button.Config{
		Text:    "Link",
		Variant: theme.VariantLink,
		Size:    theme.SizeDefault,
		OnClick: func() {
			log.Println("Link button clicked!")
		},
	})

	// Theme toggle button
	var themeToggleBtn *button.Button
	themeToggleBtn = button.New(button.Config{
		Text:    "🌙 Dark Mode",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			th.ToggleDark()
			// Update window colors to match new theme
			updateWindowColors(w, th)
			// Update button text based on new theme state
			if th.IsDark {
				themeToggleBtn.SetText("☀️ Light Mode")
			} else {
				themeToggleBtn.SetText("🌙 Dark Mode")
			}
			w.Invalidate() // Force immediate redraw
			log.Println("Theme toggled!")
		},
	})

	// Zoom control buttons
	zoomInBtn := button.New(button.Config{
		Text:    "+",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			if zoomScale < maxZoom {
				zoomScale += zoomStep
				if zoomScale > maxZoom {
					zoomScale = maxZoom
				}
				zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
				w.Invalidate()
				log.Printf("Zoom in (button): %.1f%%", zoomScale*100)
			}
		},
	})

	zoomOutBtn := button.New(button.Config{
		Text:    "-",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			if zoomScale > minZoom {
				zoomScale -= zoomStep
				if zoomScale < minZoom {
					zoomScale = minZoom
				}
				zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
				w.Invalidate()
				log.Printf("Zoom out (button): %.1f%%", zoomScale*100)
			}
		},
	})

	zoomResetBtn := button.New(button.Config{
		Text:    "Reset",
		Variant: theme.VariantOutline,
		Size:    theme.SizeSM,
		OnClick: func() {
			zoomScale = 1.0
			zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
			w.Invalidate()
			log.Printf("Reset zoom (button): %.1f%%", zoomScale*100)
		},
	})

	// Components
	demoCard := card.New(card.Config{
		Variant: theme.VariantDefault,
	})

	// Input component
	textInput := input.Text("Enter your name...")

	titleLabel := label.NewTypography("Demo-app", label.H1, "")
	subtitleLabel := label.NewTypography("A shadcn/ui port for Gio", label.P, "")
	zoomHelpLabel := label.NewTypography("Use buttons or Ctrl+/- to zoom, Ctrl+0 to reset", label.Small, "")

	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Process global keyboard events (without focus stealing)
			for {
				ev, ok := gtx.Event(
					key.Filter{
						Name:     "+",
						Required: key.ModCtrl,
					},
					key.Filter{
						Name:     "=",
						Required: key.ModCtrl,
					},
					key.Filter{
						Name:     "-",
						Required: key.ModCtrl,
					},
					key.Filter{
						Name:     "0",
						Required: key.ModCtrl,
					},
				)
				if !ok {
					break
				}

				if e, ok := ev.(key.Event); ok {
					// Debug logging for zoom key events
					if e.State == key.Press {
						log.Printf("Global key pressed: '%s' with modifiers: %v", e.Name, e.Modifiers)
					}

					// Handle zoom keyboard shortcuts
					if e.State == key.Press {
						switch {
						case (e.Name == "+" || e.Name == "=") && e.Modifiers == key.ModCtrl:
							// Zoom in (Ctrl+/= or Ctrl+numpad+)
							if zoomScale < maxZoom {
								zoomScale += zoomStep
								if zoomScale > maxZoom {
									zoomScale = maxZoom
								}
								zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
								w.Invalidate()
								log.Printf("Zoom in: %.1f%% (key: %s)", zoomScale*100, e.Name)
							}
						case e.Name == "-" && e.Modifiers == key.ModCtrl:
							// Zoom out (Ctrl- or Ctrl+numpad-)
							if zoomScale > minZoom {
								zoomScale -= zoomStep
								if zoomScale < minZoom {
									zoomScale = minZoom
								}
								zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
								w.Invalidate()
								log.Printf("Zoom out: %.1f%% (key: %s)", zoomScale*100, e.Name)
							}
						case e.Name == "0" && e.Modifiers == key.ModCtrl:
							// Reset zoom (Ctrl+0)
							zoomScale = 1.0
							zoomLabel.SetText(fmt.Sprintf("Zoom: %.0f%%", zoomScale*100))
							w.Invalidate()
							log.Printf("Reset zoom: %.1f%% (key: %s)", zoomScale*100, e.Name)
						}
					}
				}
			}

			// Apply zoom scale to the context
			if zoomScale != 1.0 {
				scaledMetric := unit.Metric{
					PxPerDp: e.Metric.PxPerDp * zoomScale,
					PxPerSp: e.Metric.PxPerSp * zoomScale,
				}
				gtx.Metric = scaledMetric
			}

			// Set background color - fill the entire window first
			background := clip.Rect{Max: gtx.Constraints.Max}.Op()
			paint.FillShape(gtx.Ops, th.Colors.Background, background)

			// Main layout
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				// Custom title bar
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return tb.Layout(gtx, th, w)
				}),

				// Header
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    th.Spacing.Space8,
						Bottom: th.Spacing.Space8,
						Left:   th.Spacing.Space8,
						Right:  th.Spacing.Space8,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:      layout.Horizontal,
							Alignment: layout.Middle,
						}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis: layout.Vertical,
								}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return titleLabel.Layout(gtx, th)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return subtitleLabel.Layout(gtx, th)
									}),
								)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis:      layout.Vertical,
									Alignment: layout.End,
								}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return themeToggleBtn.Layout(gtx, th)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Flex{
											Axis:      layout.Horizontal,
											Alignment: layout.Middle,
										}.Layout(gtx,
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return zoomOutBtn.Layout(gtx, th)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return zoomInBtn.Layout(gtx, th)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return zoomResetBtn.Layout(gtx, th)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return layout.Spacer{Width: th.Spacing.Space4}.Layout(gtx)
											}),
											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												return zoomLabel.Layout(gtx, th)
											}),
										)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return zoomHelpLabel.Layout(gtx, th)
									}),
								)
							}),
						)
					})
				}),

				// Main content
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    th.Spacing.Space4,
						Bottom: th.Spacing.Space8,
						Left:   th.Spacing.Space8,
						Right:  th.Spacing.Space8,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return demoCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{
								Axis: layout.Vertical,
							}.Layout(gtx,
								// Card header
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									cardTitle := label.NewTypography("Component Examples", label.H3, "")
									return layout.Flex{
										Axis: layout.Vertical,
									}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return cardTitle.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
										}),
									)
								}),

								// Button examples
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Buttons", label.H4, "")
									return layout.Flex{
										Axis: layout.Vertical,
									}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return sectionTitle.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layoutButtonRow(gtx, th, primaryBtn, destructiveBtn, outlineBtn)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layoutButtonRow(gtx, th, secondaryBtn, ghostBtn, linkBtn)
										}),
									)
								}),

								// Input examples
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Spacer{Height: th.Spacing.Space8}.Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Input Components", label.H4, "")
									return layout.Flex{
										Axis: layout.Vertical,
									}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return sectionTitle.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											// Create a constrained context for inputs
											maxWidth := gtx.Metric.Dp(400)
											if gtx.Constraints.Max.X < maxWidth {
												maxWidth = gtx.Constraints.Max.X
											}

											gtx.Constraints.Max.X = maxWidth
											gtx.Constraints.Min.X = maxWidth

											return textInput.Layout(gtx, th)
										}),
									)
								}),

								// Typography examples
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Spacer{Height: th.Spacing.Space8}.Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									sectionTitle := label.NewTypography("Typography", label.H4, "")
									h1Example := label.NewTypography("Heading 1", label.H1, "")
									h2Example := label.NewTypography("Heading 2", label.H2, "")
									h3Example := label.NewTypography("Heading 3", label.H3, "")
									bodyExample := label.NewTypography("This is a paragraph of body text demonstrating the typography system.", label.P, "")
									smallExample := label.NewTypography("Small text for captions and fine print.", label.Small, "")
									mutedExample := label.NewTypography("Muted text for secondary information.", label.Muted, "")

									return layout.Flex{
										Axis: layout.Vertical,
									}.Layout(gtx,
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return sectionTitle.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return h1Example.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return h2Example.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return h3Example.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return bodyExample.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return smallExample.Layout(gtx, th)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return mutedExample.Layout(gtx, th)
										}),
									)
								}),
							)
						})
					})
				}),
			)

			e.Frame(&ops)
		}
	}
}
