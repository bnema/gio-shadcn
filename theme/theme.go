package theme

import (
	"fmt"
	"image/color"
)

import (
	"gioui.org/layout"
)

// Theme represents the complete theme configuration
type Theme struct {
	Colors     ColorScheme
	DarkColors ColorScheme
	Typography Typography
	Spacing    SpacingScale
	Radius     RadiusScale
	IsDark     bool
}

// New creates a new theme with light colors by default
func New() *Theme {
	return &Theme{
		Colors:     LightColorScheme(),
		DarkColors: DarkColorScheme(),
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
		DarkColors: LightColorScheme(),
		Typography: DefaultTypography(),
		Spacing:    DefaultSpacing(),
		Radius:     DefaultRadius(),
		IsDark:     true,
	}
}

// ToggleDark switches between light and dark color schemes
func (t *Theme) ToggleDark() {
	if t.IsDark {
		// Switch to light mode - swap current colors back
		t.Colors, t.DarkColors = t.DarkColors, t.Colors
		t.IsDark = false
	} else {
		// Switch to dark mode - swap colors
		t.Colors, t.DarkColors = t.DarkColors, t.Colors
		t.IsDark = true
	}
}

// ValidateTheme validates that a theme has all required fields
func ValidateTheme(t *Theme) error {
	if t == nil {
		return fmt.Errorf("theme cannot be nil")
	}
	
	// Validate color scheme
	if err := validateColorScheme(&t.Colors); err != nil {
		return fmt.Errorf("invalid light colors: %w", err)
	}
	
	if err := validateColorScheme(&t.DarkColors); err != nil {
		return fmt.Errorf("invalid dark colors: %w", err)
	}
	
	return nil
}

func validateColorScheme(cs *ColorScheme) error {
	if cs == nil {
		return fmt.Errorf("color scheme cannot be nil")
	}
	
	// Check that all colors have non-zero alpha
	colors := []struct {
		name  string
		color color.NRGBA
	}{
		{"background", cs.Background},
		{"foreground", cs.Foreground},
		{"primary", cs.Primary},
		{"primary-foreground", cs.PrimaryFg},
		{"secondary", cs.Secondary},
		{"secondary-foreground", cs.SecondaryFg},
		{"border", cs.Border},
	}
	
	for _, c := range colors {
		if c.color.A == 0 {
			return fmt.Errorf("color %s has zero alpha", c.name)
		}
	}
	
	return nil
}

// Component represents a themeable component
type Component interface {
	Layout(gtx layout.Context, theme *Theme) layout.Dimensions
	Update(gtx layout.Context) ComponentState
}

// ComponentState represents the state of a component
type ComponentState interface {
	IsActive() bool
	IsHovered() bool
	IsPressed() bool
	IsDisabled() bool
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
