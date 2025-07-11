package theme

import (
	"gioui.org/layout"
)

// Theme represents the complete theme configuration
type Theme struct {
	Colors     ColorScheme
	Typography Typography
	Spacing    SpacingScale
	Radius     RadiusScale
	IsDark     bool
}

// New creates a new theme with light colors by default
func New() *Theme {
	return &Theme{
		Colors:     LightColorScheme(),
		Typography: DefaultTypography(),
		Spacing:    DefaultSpacing(),
		Radius:     DefaultRadius(),
		IsDark:     false,
	}
}

// NewDark creates a new theme with dark colors
func NewDark() *Theme {
	return &Theme{
		Colors:     DarkColorScheme(),
		Typography: DefaultTypography(),
		Spacing:    DefaultSpacing(),
		Radius:     DefaultRadius(),
		IsDark:     true,
	}
}

// ToggleDark switches between light and dark color schemes
func (t *Theme) ToggleDark() {
	if t.IsDark {
		t.Colors = LightColorScheme()
		t.IsDark = false
	} else {
		t.Colors = DarkColorScheme()
		t.IsDark = true
	}
}

// Component represents a themeable component
type Component interface {
	Layout(gtx layout.Context, theme *Theme) layout.Dimensions
}

// ComponentState represents the state of a component
type ComponentState interface {
	Update(gtx layout.Context) ComponentState
}

// Variant represents a component variant
type Variant string

// Size represents a component size
type Size string

// Common variants
const (
	VariantDefault     Variant = "default"
	VariantDestructive Variant = "destructive"
	VariantOutline     Variant = "outline"
	VariantSecondary   Variant = "secondary"
	VariantGhost       Variant = "ghost"
	VariantLink        Variant = "link"
)

// Common sizes
const (
	SizeDefault Size = "default"
	SizeSM      Size = "sm"
	SizeLG      Size = "lg"
	SizeIcon    Size = "icon"
)
