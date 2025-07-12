/*
Package theme provides a comprehensive theming system for gio-shadcn components.

The theme system enables consistent styling across all UI components with support
for light/dark modes, customizable color schemes, typography, spacing, and border
radius scales.

# Quick Start

Create a default light theme:

	th := theme.New()

Create a dark theme:

	th := theme.NewDark()

Toggle between light and dark modes:

	th.ToggleDark()

# Theme Structure

A theme consists of:
• Colors - Primary, secondary, background colors with light/dark variants
• Typography - Font sizes, weights, and styling
• Spacing - Consistent spacing scale for layout
• Radius - Border radius scale for rounded corners

# Component Integration

All gio-shadcn components accept a *Theme parameter in their Layout method:

	button := button.New(button.Config{Text: "Click me"})
	dims := button.Layout(gtx, th)

# Customization

Themes can be customized programmatically or loaded from JSON configuration files.
See the theme JSON format in the project's theme.json example file.
*/
package theme

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
)

// Theme represents the complete theme configuration for gio-shadcn components.
// It contains all styling information including colors, typography, spacing,.
// and border radius scales. The theme supports both light and dark color schemes
// that can be toggled at runtime.
//
// Example usage:.
//
//	th := theme.New()              // Create light theme
//	th.ToggleDark()               // Switch to dark mode
//	button.Layout(gtx, th)        // Use theme in components
type Theme struct {
	Colors     ColorScheme
	DarkColors ColorScheme
	Typography Typography
	Spacing    SpacingScale
	Radius     RadiusScale
	IsDark     bool
}

// New creates a new theme with light colors by default.
// This is the most common way to create a theme for gio-shadcn applications.
// The theme includes default color schemes for both light and dark modes,.
// typography settings, spacing scale, and border radius scale.
//
// Example:.
//
//	th := theme.New()
//	// Theme is ready to use with all components
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

// NewDark creates a new theme with dark colors by default.
// This is useful when you want to start your application in dark mode.
// The theme will have dark colors active and light colors stored for toggling.
//
// Example:.
//
//	th := theme.NewDark()
//	// Theme starts in dark mode
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

// ToggleDark switches between light and dark color schemes.
// This method swaps the current Colors with DarkColors, allowing.
// runtime theme switching. Call window.Invalidate() after toggling
// to force a UI redraw.
//
// Example:.
//
//	th := theme.New()          // Starts in light mode
//	th.ToggleDark()           // Switches to dark mode
//	th.ToggleDark()           // Switches back to light mode
//	window.Invalidate()       // Force UI refresh
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

// ValidateTheme validates that a theme has all required fields and valid colors.
// This function checks that the theme is not nil and that all color schemes.
// have valid colors with non-zero alpha values. Use this when loading themes
// from external sources or after programmatic modifications.
//
// Returns an error if validation fails, nil if the theme is valid.
//
// Example:.
//
//	th := theme.New()
//	if err := theme.ValidateTheme(th); err != nil {
//		log.Printf("Theme validation failed: %v", err)
//	}
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

// Component represents a themeable UI component in the gio-shadcn library.
// All gio-shadcn components implement this interface, ensuring consistent.
// behavior across the component system. Components must be able to layout
// themselves given a graphics context and theme, and provide state information.
type Component interface {
	Layout(gtx layout.Context, theme *Theme) layout.Dimensions
	Update(gtx layout.Context) ComponentState
}

// ComponentState represents the current interactive state of a UI component.
// This interface allows components to communicate their state to parent layouts.
// or other components, enabling conditional styling, tooltips, or other.
// state-dependent behavior.
type ComponentState interface {
	IsActive() bool
	IsHovered() bool
	IsPressed() bool
	IsDisabled() bool
}

// Variant represents a visual style variant for components.
// Components use variants to determine their appearance (colors, styling).
// Each variant corresponds to a different semantic meaning or visual emphasis.
type Variant string

// Size represents a size variant for components.
// Components use sizes to determine their dimensions and internal spacing.
// Sizes provide consistent scaling across the component system.
type Size string

// Standard component variants used across the gio-shadcn component library.
// These variants provide consistent semantic meaning and visual styling.
const (
	VariantDefault     Variant = "default"     // Primary action, solid background
	VariantDestructive Variant = "destructive" // Dangerous actions, red theme
	VariantOutline     Variant = "outline"     // Border only, transparent background
	VariantSecondary   Variant = "secondary"   // Less prominent than default
	VariantGhost       Variant = "ghost"       // Minimal styling, appears on hover
	VariantLink        Variant = "link"        // Styled like a hyperlink
)

// Standard component sizes used across the gio-shadcn component library.
// These sizes provide consistent scaling and spacing across all components.
const (
	SizeDefault Size = "default" // Standard size for most use cases
	SizeSM      Size = "sm"      // Small size for compact layouts
	SizeLG      Size = "lg"      // Large size for emphasis or accessibility
	SizeIcon    Size = "icon"    // Square size optimized for icons
)
