package theme

import (
	"image/color"
)

// Common colors.
var (
	transparent = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
)

// VariantConfig represents the configuration for a component variant.
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

// createSolidVariant creates a solid color variant configuration.
func createSolidVariant(bg, fg color.NRGBA, colors *ColorScheme) VariantConfig {
	return VariantConfig{
		Background:  bg,
		Foreground:  fg,
		Border:      bg,
		BorderWidth: 0,
		HoverBg:     darken(bg, 0.1),
		HoverFg:     fg,
		ActiveBg:    darken(bg, 0.2),
		ActiveFg:    fg,
		DisabledBg:  colors.Muted,
		DisabledFg:  colors.MutedFg,
		FocusRing:   colors.Ring,
	}
}

// createTransparentVariant creates a transparent variant with hover/active states.
func createTransparentVariant(fg color.NRGBA, colors *ColorScheme, withBorder bool) VariantConfig {
	return createTransparentVariantWithOpacity(fg, colors, withBorder, 0)
}

// createTransparentVariantWithOpacity creates a transparent variant with custom opacity.
func createTransparentVariantWithOpacity(fg color.NRGBA, colors *ColorScheme, withBorder bool, bgOpacity uint8) VariantConfig {
	config := VariantConfig{
		Background:  color.NRGBA{R: colors.Background.R, G: colors.Background.G, B: colors.Background.B, A: bgOpacity},
		Foreground:  fg,
		Border:      transparent,
		BorderWidth: 0,
		HoverBg:     colors.Accent,
		HoverFg:     colors.AccentFg,
		ActiveBg:    colors.Accent,
		ActiveFg:    colors.AccentFg,
		DisabledBg:  transparent,
		DisabledFg:  colors.MutedFg,
		FocusRing:   colors.Ring,
	}

	if withBorder {
		config.Border = colors.Border
		config.BorderWidth = 1
	}

	return config
}

// GetButtonVariant returns the color configuration for a button variant.
func GetButtonVariant(variant Variant, colors *ColorScheme) VariantConfig {
	switch variant {
	case VariantDefault:
		return createSolidVariant(colors.Primary, colors.PrimaryFg, colors)

	case VariantDestructive:
		return createSolidVariant(colors.Destructive, colors.DestructiveFg, colors)

	case VariantOutline:
		return createTransparentVariant(colors.Foreground, colors, true)

	case VariantSecondary:
		return createSolidVariant(colors.Secondary, colors.SecondaryFg, colors)

	case VariantGhost:
		return createTransparentVariant(colors.Foreground, colors, false)

	case VariantLink:
		return VariantConfig{
			Background:  transparent,
			Foreground:  colors.Primary,
			Border:      transparent,
			BorderWidth: 0,
			HoverBg:     transparent,
			HoverFg:     darken(colors.Primary, 0.1),
			ActiveBg:    transparent,
			ActiveFg:    darken(colors.Primary, 0.2),
			DisabledBg:  transparent,
			DisabledFg:  colors.MutedFg,
			FocusRing:   colors.Ring,
		}

	default:
		return GetButtonVariant(VariantDefault, colors)
	}
}

// GetCardVariant returns the color configuration for a card variant.
func GetCardVariant(variant Variant, colors *ColorScheme) VariantConfig {
	switch variant {
	case VariantDefault:
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
		return GetCardVariant(VariantDefault, colors)
	}
}

// createInputVariant creates a base input variant configuration.
func createInputVariant(bg color.NRGBA, fg color.NRGBA, border color.NRGBA, colors *ColorScheme) VariantConfig {
	return VariantConfig{
		Background:  bg,
		Foreground:  fg,
		Border:      border,
		BorderWidth: 1,
		HoverBg:     bg,
		HoverFg:     fg,
		ActiveBg:    bg,
		ActiveFg:    fg,
		DisabledBg:  colors.Muted,
		DisabledFg:  colors.MutedFg,
		FocusRing:   colors.Ring,
	}
}

// GetInputVariant returns the color configuration for an input variant.
func GetInputVariant(variant Variant, colors *ColorScheme) VariantConfig {
	switch variant {
	case VariantDefault:
		return createInputVariant(colors.Input, colors.Foreground, colors.Border, colors)

	case VariantDestructive:
		config := createInputVariant(colors.Input, colors.Foreground, colors.Destructive, colors)
		config.HoverBg = lighten(colors.Input, 0.05)
		config.ActiveBg = lighten(colors.Input, 0.1)
		config.FocusRing = colors.Destructive
		return config

	case VariantOutline:
		return createTransparentVariant(colors.Foreground, colors, true)

	case VariantSecondary:
		config := createInputVariant(colors.Secondary, colors.SecondaryFg, colors.Border, colors)
		config.HoverBg = lighten(colors.Secondary, 0.05)
		config.HoverFg = colors.SecondaryFg
		config.ActiveBg = lighten(colors.Secondary, 0.1)
		config.ActiveFg = colors.SecondaryFg
		return config

	case VariantGhost:
		return createTransparentVariant(colors.Foreground, colors, false)

	default:
		return GetInputVariant(VariantDefault, colors)
	}
}

// GetTitleBarVariant returns the color configuration for a titlebar variant.
func GetTitleBarVariant(variant Variant, colors *ColorScheme) VariantConfig {
	switch variant {
	case VariantDefault:
		// Use the background colors for default titlebar
		config := createSolidVariant(colors.Background, colors.Foreground, colors)
		config.Border = colors.Border
		config.BorderWidth = 1
		return config

	case VariantDestructive:
		return createSolidVariant(colors.Destructive, colors.DestructiveFg, colors)

	case VariantOutline:
		return createTransparentVariant(colors.Foreground, colors, true)

	case VariantSecondary:
		return createSolidVariant(colors.Secondary, colors.SecondaryFg, colors)

	case VariantGhost:
		// For titlebar ghost variant, use a subtle background opacity
		return createTransparentVariantWithOpacity(colors.Foreground, colors, false, 20)

	default:
		return GetTitleBarVariant(VariantDefault, colors)
	}
}

// Helper function to darken a color.
func darken(c color.NRGBA, factor float32) color.NRGBA {
	return color.NRGBA{
		R: uint8(float32(c.R) * (1 - factor)),
		G: uint8(float32(c.G) * (1 - factor)),
		B: uint8(float32(c.B) * (1 - factor)),
		A: c.A,
	}
}

// Helper function to lighten a color.
func lighten(c color.NRGBA, factor float32) color.NRGBA {
	return color.NRGBA{
		R: uint8(float32(c.R) + (255-float32(c.R))*factor),
		G: uint8(float32(c.G) + (255-float32(c.G))*factor),
		B: uint8(float32(c.B) + (255-float32(c.B))*factor),
		A: c.A,
	}
}
