/*
Package button provides a versatile button component for gio-shadcn applications.

The button component supports multiple visual variants, sizes, and interactive states,
following shadcn/ui design principles. It provides a clean, accessible interface
with proper keyboard and mouse interaction handling.

# Quick Start

Create a basic button:

	btn := button.New(button.Config{
		Text:    "Click me",
		Variant: theme.VariantDefault,
		OnClick: func() {
			fmt.Println("Button clicked!")
		},
	})

Use in layout:

	dims := btn.Layout(gtx, th)

# Variants

The button supports these visual variants:
• VariantDefault - Primary action button with solid background
• VariantDestructive - For dangerous actions (red theme)
• VariantOutline - Border only, transparent background
• VariantSecondary - Less prominent than default
• VariantGhost - Minimal styling, appears on hover
• VariantLink - Styled like a hyperlink

# Sizes

Available sizes:
• SizeDefault - Standard size for most use cases
• SizeSM - Small size for compact layouts
• SizeLG - Large size for emphasis
• SizeIcon - Square size optimized for icons

# Features

• Multiple visual variants following shadcn/ui design
• Consistent sizing system
• Hover and active states
• Disabled state support
• Optional icon support
• Custom CSS-style class utilities
• Accessible keyboard interaction
• Theme integration with automatic color adaptation

# Examples

Basic button:

	btn := button.New(button.Config{
		Text: "Save",
		Variant: theme.VariantDefault,
	})

Destructive action:

	deleteBtn := button.New(button.Config{
		Text: "Delete",
		Variant: theme.VariantDestructive,
		OnClick: func() { confirmDelete() },
	})

Small outline button:

	cancelBtn := button.New(button.Config{
		Text: "Cancel",
		Variant: theme.VariantOutline,
		Size: theme.SizeSM,
	})
*/
package button

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Button represents a shadcn/ui button component with full state management.
// It provides a comprehensive button implementation with support for variants,.
// sizes, icons, disabled states, and custom styling through CSS-like classes.
// The button handles all user interactions including mouse and keyboard events.
//
// Example usage:.
//
//	btn := button.New(button.Config{
//		Text: "Save",
//		Variant: theme.VariantDefault,
//		OnClick: func() { save() },
//	})
//	dims := btn.Layout(gtx, theme)
type Button struct {
	// State
	clickable *widget.Clickable

	// Configuration
	Text     string
	Variant  theme.Variant
	Size     theme.Size
	Icon     *widget.Icon
	Disabled bool
	Classes  string
	OnClick  func()

	// Cached parsed styles to avoid re-parsing on every frame
	cachedStyles     utils.StyleUtility
	cachedClasses    string
	stylesCacheValid bool
}

// Option is a functional option for configuring Button components.
type Option func(*Button)

// WithVariant sets the button variant.
func WithVariant(variant theme.Variant) Option {
	return func(b *Button) {
		b.Variant = variant
	}
}

// WithSize sets the button size.
func WithSize(size theme.Size) Option {
	return func(b *Button) {
		b.Size = size
	}
}

// WithText sets the button text.
func WithText(text string) Option {
	return func(b *Button) {
		b.Text = text
	}
}

// WithIcon sets the button icon.
func WithIcon(icon *widget.Icon) Option {
	return func(b *Button) {
		b.Icon = icon
	}
}

// WithOnClick sets the click handler.
func WithOnClick(onClick func()) Option {
	return func(b *Button) {
		b.OnClick = onClick
	}
}

// WithDisabled sets the disabled state.
func WithDisabled(disabled bool) Option {
	return func(b *Button) {
		b.Disabled = disabled
	}
}

// WithClasses sets additional CSS-like classes.
func WithClasses(classes string) Option {
	return func(b *Button) {
		b.Classes = classes
	}
}

// NewButton creates a new Button with the given options.
func NewButton(options ...Option) *Button {
	b := &Button{
		clickable: &widget.Clickable{},
		Variant:   theme.VariantDefault,
		Size:      theme.SizeDefault,
	}

	for _, option := range options {
		option(b)
	}

	return b
}

// ValidateButton validates that a button has all required fields.
func ValidateButton(b *Button) error {
	if b == nil {
		return fmt.Errorf("button cannot be nil")
	}

	if b.clickable == nil {
		return fmt.Errorf("button must have a clickable widget")
	}

	if b.Text == "" && b.Icon == nil {
		return fmt.Errorf("button must have either text or icon")
	}

	return nil
}

// SafeLayout is a wrapper around Layout that validates the button first.
func (b *Button) SafeLayout(gtx layout.Context, th *theme.Theme) (layout.Dimensions, error) {
	if err := ValidateButton(b); err != nil {
		return layout.Dimensions{}, err
	}

	if th == nil {
		return layout.Dimensions{}, fmt.Errorf("theme cannot be nil")
	}

	return b.Layout(gtx, th), nil
}

// Config represents button configuration for easy initialization.
// This struct provides a convenient way to configure all button properties.
// at once when creating a new button instance. All fields are optional
// and will use sensible defaults if not specified.
//
// Example:.
//
//	config := button.Config{
//		Text: "Submit",
//		Variant: theme.VariantDefault,
//		Size: theme.SizeDefault,
//		OnClick: func() { submitForm() },
//	}
//	btn := button.New(config)
type Config struct {
	Text     string
	Variant  theme.Variant
	Size     theme.Size
	Icon     *widget.Icon
	Disabled bool
	Classes  string
	OnClick  func()
}

// New creates a new button with the given configuration.
// This is the recommended way to create button instances. It initializes
// all internal state and applies the provided configuration. If no variant
// or size is specified, sensible defaults will be used.
//
// Example:.
//
//	btn := button.New(button.Config{
//		Text: "Click me",
//		Variant: theme.VariantDefault,
//		OnClick: func() { fmt.Println("Clicked!") },
//	})
func New(config Config) *Button {
	return &Button{
		clickable: new(widget.Clickable),
		Text:      config.Text,
		Variant:   config.Variant,
		Size:      config.Size,
		Icon:      config.Icon,
		Disabled:  config.Disabled,
		Classes:   config.Classes,
		OnClick:   config.OnClick,
	}
}

// Layout renders the button with the given graphics context and theme.
// This method handles all the visual rendering of the button including.
// background, borders, text, icons, and interactive states (hover, active, disabled).
// It also processes click events and calls the OnClick handler when appropriate.
//
// The layout automatically adapts to the current theme and applies the configured.
// variant and size styling. Custom CSS-like classes are also applied if specified.
//
// Returns the dimensions occupied by the button after rendering.
func (b *Button) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	// Handle click events
	if b.clickable.Clicked(gtx) && !b.Disabled && b.OnClick != nil {
		b.OnClick()
	}

	// Get variant configuration
	variant := theme.GetButtonVariant(b.Variant, &th.Colors)

	// Get size configuration
	padding, minHeight, fontSize := b.getSizeConfig(th)

	// Parse additional classes (with caching)
	var styles utils.StyleUtility
	if !b.stylesCacheValid || b.cachedClasses != b.Classes {
		styles = utils.ParseClasses(b.Classes)
		b.cachedStyles = styles
		b.cachedClasses = b.Classes
		b.stylesCacheValid = true
	} else {
		styles = b.cachedStyles
	}

	// Apply custom padding if specified
	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}

	// Determine current state colors
	bgColor := variant.Background
	fgColor := variant.Foreground

	switch {
	case b.Disabled:
		bgColor = variant.DisabledBg
		fgColor = variant.DisabledFg
	case b.clickable.Pressed():
		bgColor = variant.ActiveBg
		fgColor = variant.ActiveFg
	case b.clickable.Hovered():
		bgColor = variant.HoverBg
		fgColor = variant.HoverFg
	}

	// Apply custom background if specified
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	return b.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return b.drawButton(gtx, th, bgColor, fgColor, variant, padding, minHeight, fontSize, styles)
	})
}

// Update returns the current state of the button component.
// This method provides access to the button's interactive state including.
// whether it's currently hovered, pressed, active, or disabled. This information
// can be used by parent components for conditional rendering or behavior.
//
// Example:.
//
//	state := btn.Update(gtx)
//	if state.IsHovered() {
//		// Show tooltip or additional UI
//	}
func (b *Button) Update(gtx layout.Context) theme.ComponentState {
	return &State{
		active:   b.clickable.Clicked(gtx),
		hovered:  b.clickable.Hovered(),
		pressed:  b.clickable.Pressed(),
		disabled: b.Disabled,
	}
}

// State implements theme.ComponentState for button components.
// It tracks the current interactive state of the button including hover,.
// press, active, and disabled states. This state information can be used
// by parent components or for conditional styling.
type State struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

// IsActive returns true if the button is in an active state.
func (bs *State) IsActive() bool {
	return bs.active
}

// IsHovered returns true if the button is being hovered over.
func (bs *State) IsHovered() bool {
	return bs.hovered
}

// IsPressed returns true if the button is being pressed.
func (bs *State) IsPressed() bool {
	return bs.pressed
}

// IsDisabled returns true if the button is disabled.
func (bs *State) IsDisabled() bool {
	return bs.disabled
}

func (b *Button) drawButton(gtx layout.Context, th *theme.Theme, bgColor, fgColor color.NRGBA, variant theme.VariantConfig, padding layout.Inset, minHeight unit.Dp, fontSize unit.Sp, styles utils.StyleUtility) layout.Dimensions {
	// Create rounded rectangle clip
	radius := th.Radius.RadiusMD
	if styles.Radius > 0 {
		radius = styles.Radius
	}

	// Calculate content dimensions first
	contentDims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return b.layoutContent(gtx, th, fgColor, fontSize)
	})

	// Ensure minimum height
	finalHeight := contentDims.Size.Y
	if finalHeight < gtx.Dp(minHeight) {
		finalHeight = gtx.Dp(minHeight)
	}

	// Set constraints for stack to ensure both layers are the same size
	gtx.Constraints = layout.Exact(image.Pt(contentDims.Size.X, finalHeight))

	return layout.Stack{}.Layout(gtx,
		// Background
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			// Draw background
			rect := image.Rectangle{Max: gtx.Constraints.Min}
			rr := clip.UniformRRect(rect, gtx.Dp(radius))
			paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

			// Draw border if specified
			if variant.BorderWidth > 0 {
				border := clip.Stroke{
					Path:  rr.Path(gtx.Ops),
					Width: variant.BorderWidth,
				}
				paint.FillShape(gtx.Ops, variant.Border, border.Op())
			}

			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),

		// Content - centered vertically and horizontally
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return b.layoutContent(gtx, th, fgColor, fontSize)
				})
			})
		}),
	)
}

func (b *Button) layoutContent(gtx layout.Context, th *theme.Theme, fgColor color.NRGBA, fontSize unit.Sp) layout.Dimensions {
	switch {
	case b.Icon != nil && b.Text != "":
		// Icon + text layout
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return b.Icon.Layout(gtx, th.Colors.Foreground)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: th.Spacing.Space2}.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return b.layoutText(gtx, th, fgColor, fontSize)
			}),
		)
	case b.Icon != nil:
		// Icon only
		return b.Icon.Layout(gtx, fgColor)
	default:
		// Text only
		return b.layoutText(gtx, th, fgColor, fontSize)
	}
}

func (b *Button) layoutText(gtx layout.Context, _ *theme.Theme, fgColor color.NRGBA, fontSize unit.Sp) layout.Dimensions {
	label := material.Label(material.NewTheme(), fontSize, b.Text)
	label.Color = fgColor
	label.Alignment = text.Middle
	return label.Layout(gtx)
}

func (b *Button) getSizeConfig(th *theme.Theme) (layout.Inset, unit.Dp, unit.Sp) {
	switch b.Size {
	case theme.SizeSM:
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space3,
			Right:  th.Spacing.Space3,
		}, unit.Dp(32), th.Typography.FontSizeSM

	case theme.SizeLG:
		return layout.Inset{
			Top:    th.Spacing.Space3,
			Bottom: th.Spacing.Space3,
			Left:   th.Spacing.Space8,
			Right:  th.Spacing.Space8,
		}, unit.Dp(44), th.Typography.FontSizeBase

	case theme.SizeIcon:
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space2,
			Right:  th.Spacing.Space2,
		}, unit.Dp(36), th.Typography.FontSizeSM

	default: // SizeDefault
		return layout.Inset{
			Top:    th.Spacing.Space2,
			Bottom: th.Spacing.Space2,
			Left:   th.Spacing.Space4,
			Right:  th.Spacing.Space4,
		}, unit.Dp(36), th.Typography.FontSizeSM
	}
}

// Clicked returns true if the button was clicked.
func (b *Button) Clicked(gtx layout.Context) bool {
	return b.clickable.Clicked(gtx) && !b.Disabled
}

// SetDisabled sets the disabled state of the button.
func (b *Button) SetDisabled(disabled bool) {
	b.Disabled = disabled
}

// SetText sets the button text.
func (b *Button) SetText(text string) {
	b.Text = text
}

// SetVariant sets the button variant.
func (b *Button) SetVariant(variant theme.Variant) {
	b.Variant = variant
}

// SetOnClick sets the click handler.
func (b *Button) SetOnClick(onClick func()) {
	b.OnClick = onClick
}
