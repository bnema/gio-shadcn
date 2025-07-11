package button

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
	"image"
	"image/color"
)

// Button represents a shadcn/ui button component
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

// Config represents button configuration
type Config struct {
	Text     string
	Variant  theme.Variant
	Size     theme.Size
	Icon     *widget.Icon
	Disabled bool
	Classes  string
	OnClick  func()
}

// New creates a new button with the given configuration
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

// Layout renders the button
func (b *Button) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	// Handle click events
	if b.clickable.Clicked(gtx) && !b.Disabled && b.OnClick != nil {
		b.OnClick()
	}

	// Get variant configuration
	variant := utils.GetButtonVariant(b.Variant, &th.Colors)

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

	if b.Disabled {
		bgColor = variant.DisabledBg
		fgColor = variant.DisabledFg
	} else if b.clickable.Pressed() {
		bgColor = variant.ActiveBg
		fgColor = variant.ActiveFg
	} else if b.clickable.Hovered() {
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

func (b *Button) drawButton(gtx layout.Context, th *theme.Theme, bgColor, fgColor color.NRGBA, variant utils.VariantConfig, padding layout.Inset, minHeight unit.Dp, fontSize unit.Sp, styles utils.StyleUtility) layout.Dimensions {
	// Create rounded rectangle clip
	radius := th.Radius.RadiusMD
	if styles.Radius > 0 {
		radius = styles.Radius
	}

	// Calculate dimensions
	return layout.Stack{}.Layout(gtx,
		// Background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			// Calculate button dimensions
			dims := padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return b.layoutContent(gtx, th, fgColor, fontSize)
			})

			// Ensure minimum height
			if dims.Size.Y < gtx.Dp(minHeight) {
				dims.Size.Y = gtx.Dp(minHeight)
			}

			// Draw background
			rect := image.Rectangle{Max: dims.Size}
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

			return dims
		}),

		// Content
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				dims := b.layoutContent(gtx, th, fgColor, fontSize)

				// Ensure minimum height
				if dims.Size.Y < gtx.Dp(minHeight) {
					dims.Size.Y = gtx.Dp(minHeight)
				}

				return dims
			})
		}),
	)
}

func (b *Button) layoutContent(gtx layout.Context, th *theme.Theme, fgColor color.NRGBA, fontSize unit.Sp) layout.Dimensions {
	if b.Icon != nil && b.Text != "" {
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
	} else if b.Icon != nil {
		// Icon only
		return b.Icon.Layout(gtx, fgColor)
	} else {
		// Text only
		return b.layoutText(gtx, th, fgColor, fontSize)
	}
}

func (b *Button) layoutText(gtx layout.Context, th *theme.Theme, fgColor color.NRGBA, fontSize unit.Sp) layout.Dimensions {
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

// Clicked returns true if the button was clicked
func (b *Button) Clicked(gtx layout.Context) bool {
	return b.clickable.Clicked(gtx) && !b.Disabled
}

// SetDisabled sets the disabled state of the button
func (b *Button) SetDisabled(disabled bool) {
	b.Disabled = disabled
}

// SetText sets the button text
func (b *Button) SetText(text string) {
	b.Text = text
}

// SetVariant sets the button variant
func (b *Button) SetVariant(variant theme.Variant) {
	b.Variant = variant
}

// SetOnClick sets the click handler
func (b *Button) SetOnClick(onClick func()) {
	b.OnClick = onClick
}
