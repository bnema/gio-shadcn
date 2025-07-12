/*
Package titlebar provides a custom window titlebar component for gio-shadcn applications.

The titlebar component creates a custom, themeable window title bar for frameless
windows. It includes window controls (minimize, maximize, close), drag functionality,
and seamless integration with the theme system. This allows applications to have
a consistent, branded appearance across different operating systems.

# Quick Start

Create a basic titlebar:

	titlebar := titlebar.NewTitleBar(
		titlebar.WithTitle("My Application"),
		titlebar.WithWindow(window),
	)

Use in layout:

	dims := titlebar.Layout(gtx, th, window)

Create titlebar with custom close handler:

	titlebar := titlebar.NewTitleBar(
		titlebar.WithTitle("My App"),
		titlebar.WithWindow(window),
		titlebar.WithCloseHandler(func() {
			// Custom close logic
			if hasUnsavedChanges() {
				showSaveDialog()
			} else {
				os.Exit(0)
			}
		}),
	)

# Features

• Custom window title bar with native-like appearance
• Window controls (minimize, maximize, close) with hover states
• Drag-to-move functionality
• Custom close handler support for save dialogs
• Theme integration with proper colors and typography
• Cross-platform consistent appearance
• Maximize/restore toggle functionality
• Proper window state management

# Window Integration

The titlebar requires a reference to the window for proper functionality:

	w := app.NewWindow(
		app.Title("My App"),
		app.Decorated(false), // Disable system titlebar
	)

	titlebar := titlebar.NewTitleBar(
		titlebar.WithTitle("My App"),
		titlebar.WithWindow(w),
	)

# Window Controls

The titlebar includes three standard window controls:
• Minimize - Minimizes the window to taskbar/dock
• Maximize/Restore - Toggles between maximized and normal window state
• Close - Closes the window (supports custom handlers)

# Examples

Basic frameless window with titlebar:

	func createWindow() {
		w := app.NewWindow(
			app.Title("Custom App"),
			app.Decorated(false),
		)

		titlebar := titlebar.NewTitleBar(
			titlebar.WithTitle("Custom App"),
			titlebar.WithWindow(w),
		)

		// In your layout function:
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return titlebar.Layout(gtx, th, w)
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				// Main application content
				return mainContent.Layout(gtx)
			}),
		)
	}

Titlebar with save-on-exit confirmation:

	titlebar := titlebar.NewTitleBar(
		titlebar.WithTitle("Text Editor"),
		titlebar.WithWindow(window),
		titlebar.WithCloseHandler(func() {
			if documentModified {
				// Show save dialog
				showSaveConfirmation(func(save bool) {
					if save {
						saveDocument()
					}
					os.Exit(0)
				})
			} else {
				os.Exit(0)
			}
		}),
	)
*/
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

// TitleBar represents a custom window title bar.
type TitleBar struct {
	Title       string
	window      *interface{} // Will be set to *app.Window
	minimizeBtn *button.Button
	maximizeBtn *button.Button
	closeBtn    *button.Button
	isMaximized bool
	variant     theme.Variant
}

// Option is a functional option for configuring TitleBar components.
type Option func(*TitleBar)

// WithTitle sets the titlebar title.
func WithTitle(title string) Option {
	return func(tb *TitleBar) {
		tb.Title = title
	}
}

// WithWindow sets the window reference.
func WithWindow(window interface{}) Option {
	return func(tb *TitleBar) {
		tb.window = &window
	}
}

// WithCloseHandler sets the close handler.
func WithCloseHandler(onClose func()) Option {
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

// WithVariant sets the titlebar variant.
func WithVariant(variant theme.Variant) Option {
	return func(tb *TitleBar) {
		tb.variant = variant
	}
}

// NewTitleBar creates a new TitleBar with the given options.
func NewTitleBar(options ...Option) *TitleBar {
	tb := &TitleBar{
		variant: theme.VariantDefault,
	}

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

// Layout renders the title bar.
func (tb *TitleBar) Layout(gtx layout.Context, th *theme.Theme, _ interface{}) layout.Dimensions {
	// Set fixed height for title bar
	height := gtx.Dp(40)

	// Constrain the height
	gtx.Constraints.Min.Y = height
	gtx.Constraints.Max.Y = height

	// Get variant configuration
	variantConfig := theme.GetTitleBarVariant(tb.variant, &th.Colors)

	// Background
	paint.FillShape(gtx.Ops, variantConfig.Background, clip.Rect{Max: gtx.Constraints.Max}.Op())

	// Draw bottom border if specified
	if variantConfig.BorderWidth > 0 {
		borderHeight := int(variantConfig.BorderWidth)
		borderRect := clip.Rect{
			Min: image.Pt(0, gtx.Constraints.Max.Y-borderHeight),
			Max: gtx.Constraints.Max,
		}.Op()
		paint.FillShape(gtx.Ops, variantConfig.Border, borderRect)
	}

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

					// Center the title vertically within the titlebar
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Left: th.Spacing.Space4,
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							// Create bold title label with variant foreground color
							titleLabel := label.NewLabel(
								label.WithLabelText(tb.Title),
								label.WithTextStyle(theme.TextStyle{
									Size:   th.Typography.FontSizeSM,
									Weight: font.Bold,
									Color: &theme.ColorScheme{
										Foreground: variantConfig.Foreground,
									},
								}),
							)
							return titleLabel.Layout(gtx, th)
						})
					})
				}),

				// Window controls
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// Center the buttons vertically within the titlebar
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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
					})
				}),
			)
		}),
	)
}

// Update returns the component state for titlebar buttons.
func (tb *TitleBar) Update(gtx layout.Context) theme.ComponentState {
	return &State{
		minimizeHovered: tb.minimizeBtn.Update(gtx).IsHovered(),
		maximizeHovered: tb.maximizeBtn.Update(gtx).IsHovered(),
		closeHovered:    tb.closeBtn.Update(gtx).IsHovered(),
		active:          false,
		disabled:        false,
	}
}

// State implements ComponentState for Titlebar.
type State struct {
	minimizeHovered bool
	maximizeHovered bool
	closeHovered    bool
	active          bool
	disabled        bool
}

// IsActive returns true if the titlebar is active.
func (ts *State) IsActive() bool {
	return ts.active
}

// IsHovered returns true if any titlebar control is being hovered over.
func (ts *State) IsHovered() bool {
	return ts.minimizeHovered || ts.maximizeHovered || ts.closeHovered
}

// IsPressed returns true if the titlebar is being pressed (always false).
func (ts *State) IsPressed() bool {
	return false // Titlebar itself is not pressable
}

// IsDisabled returns true if the titlebar is disabled.
func (ts *State) IsDisabled() bool {
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

// SetTitle updates the title bar title.
func (tb *TitleBar) SetTitle(title string) {
	tb.Title = title
}
