package theme

import (
	"image/color"
)

// ColorScheme represents the color palette for the theme
type ColorScheme struct {
	// Core colors
	Background color.NRGBA // --background
	Foreground color.NRGBA // --foreground
	Card       color.NRGBA // --card
	CardFg     color.NRGBA // --card-foreground
	Popover    color.NRGBA // --popover
	PopoverFg  color.NRGBA // --popover-foreground

	// Primary colors
	Primary   color.NRGBA // --primary
	PrimaryFg color.NRGBA // --primary-foreground

	// Secondary colors
	Secondary   color.NRGBA // --secondary
	SecondaryFg color.NRGBA // --secondary-foreground

	// Muted colors
	Muted   color.NRGBA // --muted
	MutedFg color.NRGBA // --muted-foreground

	// Accent colors
	Accent   color.NRGBA // --accent
	AccentFg color.NRGBA // --accent-foreground

	// State colors
	Destructive   color.NRGBA // --destructive
	DestructiveFg color.NRGBA // --destructive-foreground

	// Border colors
	Border color.NRGBA // --border
	Input  color.NRGBA // --input
	Ring   color.NRGBA // --ring
}

// LightColorScheme returns the default light color scheme
func LightColorScheme() ColorScheme {
	return ColorScheme{
		Background:    color.NRGBA{R: 255, G: 255, B: 255, A: 255}, // white
		Foreground:    color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		Card:          color.NRGBA{R: 255, G: 255, B: 255, A: 255}, // white
		CardFg:        color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		Popover:       color.NRGBA{R: 255, G: 255, B: 255, A: 255}, // white
		PopoverFg:     color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		Primary:       color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		PrimaryFg:     color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		Secondary:     color.NRGBA{R: 244, G: 244, B: 245, A: 255}, // zinc-100
		SecondaryFg:   color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		Muted:         color.NRGBA{R: 244, G: 244, B: 245, A: 255}, // zinc-100
		MutedFg:       color.NRGBA{R: 82, G: 82, B: 91, A: 255},    // zinc-500
		Accent:        color.NRGBA{R: 244, G: 244, B: 245, A: 255}, // zinc-100
		AccentFg:      color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		Destructive:   color.NRGBA{R: 239, G: 68, B: 68, A: 255},   // red-500
		DestructiveFg: color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		Border:        color.NRGBA{R: 228, G: 228, B: 231, A: 255}, // zinc-200
		Input:         color.NRGBA{R: 228, G: 228, B: 231, A: 255}, // zinc-200
		Ring:          color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
	}
}

// DarkColorScheme returns the default dark color scheme
func DarkColorScheme() ColorScheme {
	return ColorScheme{
		Background:    color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		Foreground:    color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		Card:          color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		CardFg:        color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		Popover:       color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		PopoverFg:     color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		Primary:       color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		PrimaryFg:     color.NRGBA{R: 9, G: 9, B: 11, A: 255},      // zinc-950
		Secondary:     color.NRGBA{R: 39, G: 39, B: 42, A: 255},    // zinc-800
		SecondaryFg:   color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		Muted:         color.NRGBA{R: 39, G: 39, B: 42, A: 255},    // zinc-800
		MutedFg:       color.NRGBA{R: 161, G: 161, B: 170, A: 255}, // zinc-400
		Accent:        color.NRGBA{R: 39, G: 39, B: 42, A: 255},    // zinc-800
		AccentFg:      color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		Destructive:   color.NRGBA{R: 127, G: 29, B: 29, A: 255},   // red-900
		DestructiveFg: color.NRGBA{R: 250, G: 250, B: 250, A: 255}, // zinc-50
		Border:        color.NRGBA{R: 39, G: 39, B: 42, A: 255},    // zinc-800
		Input:         color.NRGBA{R: 39, G: 39, B: 42, A: 255},    // zinc-800
		Ring:          color.NRGBA{R: 212, G: 212, B: 216, A: 255}, // zinc-300
	}
}
