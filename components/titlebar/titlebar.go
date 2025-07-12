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

// TitleBarOption is a functional option for configuring TitleBar components
type TitleBarOption func(*TitleBar)

// WithTitle sets the titlebar title
func WithTitle(title string) TitleBarOption {
	return func(tb *TitleBar) {
		tb.Title = title
	}
}

// WithWindow sets the window reference
func WithWindow(window interface{}) TitleBarOption {
	return func(tb *TitleBar) {
		tb.window = &window
	}
}

// WithCloseHandler sets the close handler
func WithCloseHandler(onClose func()) TitleBarOption {
	return func(tb *TitleBar) {
		// Update the close button with the custom handler
		tb.closeBtn = button.NewButton(
			button.WithText("✕"),
			button.WithVariant(theme.VariantGhost),
			button.WithSize(theme.SizeSM),
			button.WithOnClick(func() {
				if onClose != nil {
					onClose()
				} else {
					tb.close()
				}
			}),
		)
	}
}

// NewTitleBar creates a new TitleBar with the given options
func NewTitleBar(options ...TitleBarOption) *TitleBar {
	tb := &TitleBar{}

	// Initialize default window control buttons
	tb.minimizeBtn = button.NewButton(
		button.WithText("−"),
		button.WithVariant(theme.VariantGhost),
		button.WithSize(theme.SizeSM),
		button.WithOnClick(func() {
			tb.minimize()
		}),
	)

	tb.maximizeBtn = button.NewButton(
		button.WithText("☐"),
		button.WithVariant(theme.VariantGhost),
		button.WithSize(theme.SizeSM),
		button.WithOnClick(func() {
			tb.toggleMaximize()
		}),
	)

	tb.closeBtn = button.NewButton(
		button.WithText("✕"),
		button.WithVariant(theme.VariantGhost),
		button.WithSize(theme.SizeSM),
		button.WithOnClick(func() {
			tb.close()
		}),
	)

	// Apply options
	for _, option := range options {
		option(tb)
	}

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
								titleLabel := label.NewLabel(
									label.WithLabelText(tb.Title),
									label.WithTextStyle(theme.TextStyle{
										Size:   th.Typography.FontSizeSM,
										Weight: font.Bold,
										Color:  &th.Colors,
									}),
								)
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

// Update returns the component state for titlebar buttons
func (tb *TitleBar) Update(gtx layout.Context) theme.ComponentState {
	return &TitlebarState{
		minimizeHovered: tb.minimizeBtn.Update(gtx).IsHovered(),
		maximizeHovered: tb.maximizeBtn.Update(gtx).IsHovered(),
		closeHovered:    tb.closeBtn.Update(gtx).IsHovered(),
		active:          false,
		disabled:        false,
	}
}

// TitlebarState implements ComponentState for Titlebar
type TitlebarState struct {
	minimizeHovered bool
	maximizeHovered bool
	closeHovered    bool
	active          bool
	disabled        bool
}

func (ts *TitlebarState) IsActive() bool {
	return ts.active
}

func (ts *TitlebarState) IsHovered() bool {
	return ts.minimizeHovered || ts.maximizeHovered || ts.closeHovered
}

func (ts *TitlebarState) IsPressed() bool {
	return false // Titlebar itself is not pressable
}

func (ts *TitlebarState) IsDisabled() bool {
	return ts.disabled
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
