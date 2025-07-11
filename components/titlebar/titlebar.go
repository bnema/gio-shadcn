package titlebar

import (
	"image"

	"gioui.org/font"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/bnema/gio-shadcn/components/button"
	"github.com/bnema/gio-shadcn/components/label"
	"github.com/bnema/gio-shadcn/theme"
)

// TitleBar represents a custom window title bar
type TitleBar struct {
	Title       string
	window      *interface{} // Will be set to *app.Window
	minimizeBtn *button.Button
	maximizeBtn *button.Button
	closeBtn    *button.Button
	isMaximized bool
}

// Config holds the configuration for a TitleBar
type Config struct {
	Title   string
	Window  interface{} // *app.Window
	OnClose func()
}

// New creates a new TitleBar instance
func New(cfg Config) *TitleBar {
	tb := &TitleBar{
		Title:  cfg.Title,
		window: &cfg.Window,
	}

	// Initialize window control buttons
	tb.minimizeBtn = button.New(button.Config{
		Text:    "−",
		Variant: theme.VariantGhost,
		Size:    theme.SizeSM,
		OnClick: func() {
			tb.minimize()
		},
	})

	tb.maximizeBtn = button.New(button.Config{
		Text:    "☐",
		Variant: theme.VariantGhost,
		Size:    theme.SizeSM,
		OnClick: func() {
			tb.toggleMaximize()
		},
	})

	tb.closeBtn = button.New(button.Config{
		Text:    "✕",
		Variant: theme.VariantGhost,
		Size:    theme.SizeSM,
		OnClick: func() {
			if cfg.OnClose != nil {
				cfg.OnClose()
			} else {
				tb.close()
			}
		},
	})

	return tb
}

// Layout renders the title bar
func (tb *TitleBar) Layout(gtx layout.Context, th *theme.Theme, w interface{}) layout.Dimensions {
	// Set fixed height for title bar
	height := gtx.Dp(40)

	// Constrain the height
	gtx.Constraints.Min.Y = height
	gtx.Constraints.Max.Y = height

	// Background
	paint.FillShape(gtx.Ops, th.Colors.Background, clip.Rect{Max: gtx.Constraints.Max}.Op())

	// Draw bottom border
	borderHeight := gtx.Dp(1)
	borderRect := clip.Rect{
		Min: image.Pt(0, gtx.Constraints.Max.Y-borderHeight),
		Max: gtx.Constraints.Max,
	}.Op()
	paint.FillShape(gtx.Ops, th.Colors.Border, borderRect)

	// Layout content with explicit height constraint
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			// Ensure the content fills the full height
			gtx.Constraints.Min = gtx.Constraints.Max
			return layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
			}.Layout(gtx,
				// Draggable area with title
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					// Register the drag area for system move action
					defer clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops).Pop()
					system.ActionInputOp(system.ActionMove).Add(gtx.Ops)

					return layout.Flex{
						Axis:      layout.Horizontal,
						Alignment: layout.Middle,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left: th.Spacing.Space4,
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								// Create bold title label
								titleLabel := label.New(label.Config{
									Text: tb.Title,
									TextStyle: theme.TextStyle{
										Size:   th.Typography.FontSizeSM,
										Weight: font.Bold,
										Color:  &th.Colors,
									},
								})
								return titleLabel.Layout(gtx, th)
							})
						}),
					)
				}),

				// Window controls
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return tb.minimizeBtn.Layout(gtx, th)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return tb.maximizeBtn.Layout(gtx, th)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return tb.closeBtn.Layout(gtx, th)
						}),
					)
				}),
			)
		}),
	)
}

func (tb *TitleBar) minimize() {
	// Platform-specific minimize implementation
	// This requires access to the window instance
	if window, ok := (*tb.window).(interface{ Perform(system.Action) }); ok {
		window.Perform(system.ActionMinimize)
	}
}

func (tb *TitleBar) toggleMaximize() {
	// Toggle between maximize and restore
	if window, ok := (*tb.window).(interface{ Perform(system.Action) }); ok {
		if tb.isMaximized {
			window.Perform(system.ActionUnmaximize)
			tb.maximizeBtn.SetText("☐")
		} else {
			window.Perform(system.ActionMaximize)
			tb.maximizeBtn.SetText("❐")
		}
		tb.isMaximized = !tb.isMaximized
	}
}

func (tb *TitleBar) close() {
	// Close the window
	if window, ok := (*tb.window).(interface{ Perform(system.Action) }); ok {
		window.Perform(system.ActionClose)
	}
}

// SetTitle updates the title bar title
func (tb *TitleBar) SetTitle(title string) {
	tb.Title = title
}
