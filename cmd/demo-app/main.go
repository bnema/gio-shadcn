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

// updateWindowColors updates the window colors based on the current theme
func updateWindowColors(w *app.Window, th *theme.Theme) {
	w.Option(app.NavigationColor(th.Colors.Background))
	w.Option(app.StatusColor(th.Colors.Background))
}

func run(w *app.Window) error {
	// Initialize theme
	th := theme.New()

	// Set initial window colors to match theme
	updateWindowColors(w, th)

	// Initialize title bar
	tb := titlebar.New(titlebar.Config{
		Title:  "Gio-shadcn Demo",
		Window: w,
	})

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

	titleLabel := label.NewTypography("Welcome to Gio-shadcn", label.H1, "")
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
											return layout.Flex{
												Axis: layout.Horizontal,
											}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return primaryBtn.Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return destructiveBtn.Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return outlineBtn.Layout(gtx, th)
												}),
											)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Spacer{Height: th.Spacing.Space4}.Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Flex{
												Axis: layout.Horizontal,
											}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return secondaryBtn.Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return ghostBtn.Layout(gtx, th)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return linkBtn.Layout(gtx, th)
												}),
											)
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
