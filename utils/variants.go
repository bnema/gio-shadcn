package utils

import (
	"github.com/bnema/gio-shadcn/theme"
	"image/color"
)

// VariantConfig represents the configuration for a component variant
type VariantConfig struct {
	Background  color.NRGBA
	Foreground  color.NRGBA
	Border      color.NRGBA
	BorderWidth float32
	HoverBg     color.NRGBA
	HoverFg     color.NRGBA
	ActiveBg    color.NRGBA
	ActiveFg    color.NRGBA
	DisabledBg  color.NRGBA
	DisabledFg  color.NRGBA
	FocusRing   color.NRGBA
}

// GetButtonVariant returns the color configuration for a button variant
func GetButtonVariant(variant theme.Variant, colors *theme.ColorScheme) VariantConfig {
	switch variant {
	case theme.VariantDefault:
		return VariantConfig{
			Background:  colors.Primary,
			Foreground:  colors.PrimaryFg,
			Border:      colors.Primary,
			BorderWidth: 0,
			HoverBg:     darken(colors.Primary, 0.1),
			HoverFg:     colors.PrimaryFg,
			ActiveBg:    darken(colors.Primary, 0.2),
			ActiveFg:    colors.PrimaryFg,
			DisabledBg:  colors.Muted,
			DisabledFg:  colors.MutedFg,
			FocusRing:   colors.Ring,
		}

	case theme.VariantDestructive:
		return VariantConfig{
			Background:  colors.Destructive,
			Foreground:  colors.DestructiveFg,
			Border:      colors.Destructive,
			BorderWidth: 0,
			HoverBg:     darken(colors.Destructive, 0.1),
			HoverFg:     colors.DestructiveFg,
			ActiveBg:    darken(colors.Destructive, 0.2),
			ActiveFg:    colors.DestructiveFg,
			DisabledBg:  colors.Muted,
			DisabledFg:  colors.MutedFg,
			FocusRing:   colors.Ring,
		}

	case theme.VariantOutline:
		return VariantConfig{
			Background:  color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			Foreground:  colors.Foreground,
			Border:      colors.Border,
			BorderWidth: 1,
			HoverBg:     colors.Accent,
			HoverFg:     colors.AccentFg,
			ActiveBg:    colors.Accent,
			ActiveFg:    colors.AccentFg,
			DisabledBg:  color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			DisabledFg:  colors.MutedFg,
			FocusRing:   colors.Ring,
		}

	case theme.VariantSecondary:
		return VariantConfig{
			Background:  colors.Secondary,
			Foreground:  colors.SecondaryFg,
			Border:      colors.Secondary,
			BorderWidth: 0,
			HoverBg:     darken(colors.Secondary, 0.1),
			HoverFg:     colors.SecondaryFg,
			ActiveBg:    darken(colors.Secondary, 0.2),
			ActiveFg:    colors.SecondaryFg,
			DisabledBg:  colors.Muted,
			DisabledFg:  colors.MutedFg,
			FocusRing:   colors.Ring,
		}

	case theme.VariantGhost:
		return VariantConfig{
			Background:  color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			Foreground:  colors.Foreground,
			Border:      color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			BorderWidth: 0,
			HoverBg:     colors.Accent,
			HoverFg:     colors.AccentFg,
			ActiveBg:    colors.Accent,
			ActiveFg:    colors.AccentFg,
			DisabledBg:  color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			DisabledFg:  colors.MutedFg,
			FocusRing:   colors.Ring,
		}

	case theme.VariantLink:
		return VariantConfig{
			Background:  color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			Foreground:  colors.Primary,
			Border:      color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			BorderWidth: 0,
			HoverBg:     color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			HoverFg:     darken(colors.Primary, 0.1),
			ActiveBg:    color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			ActiveFg:    darken(colors.Primary, 0.2),
			DisabledBg:  color.NRGBA{R: 0, G: 0, B: 0, A: 0}, // transparent
			DisabledFg:  colors.MutedFg,
			FocusRing:   colors.Ring,
		}

	default:
		return GetButtonVariant(theme.VariantDefault, colors)
	}
}

// GetCardVariant returns the color configuration for a card variant
func GetCardVariant(variant theme.Variant, colors *theme.ColorScheme) VariantConfig {
	switch variant {
	case theme.VariantDefault:
		return VariantConfig{
			Background:  colors.Card,
			Foreground:  colors.CardFg,
			Border:      colors.Border,
			BorderWidth: 1,
			HoverBg:     colors.Card,
			HoverFg:     colors.CardFg,
			ActiveBg:    colors.Card,
			ActiveFg:    colors.CardFg,
			DisabledBg:  colors.Muted,
			DisabledFg:  colors.MutedFg,
			FocusRing:   colors.Ring,
		}

	default:
		return GetCardVariant(theme.VariantDefault, colors)
	}
}

// Helper function to darken a color
func darken(c color.NRGBA, factor float32) color.NRGBA {
	return color.NRGBA{
		R: uint8(float32(c.R) * (1 - factor)),
		G: uint8(float32(c.G) * (1 - factor)),
		B: uint8(float32(c.B) * (1 - factor)),
		A: c.A,
	}
}

// Helper function to lighten a color
func lighten(c color.NRGBA, factor float32) color.NRGBA {
	return color.NRGBA{
		R: uint8(float32(c.R) + (255-float32(c.R))*factor),
		G: uint8(float32(c.G) + (255-float32(c.G))*factor),
		B: uint8(float32(c.B) + (255-float32(c.B))*factor),
		A: c.A,
	}
}

// Helper function to adjust alpha
func withAlpha(c color.NRGBA, alpha uint8) color.NRGBA {
	return color.NRGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: alpha,
	}
}
