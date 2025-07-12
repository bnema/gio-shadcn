/*
Package utils provides utility functions and CSS-like class parsing for gio-shadcn components.

This package contains helper functions for styling, class name management, and
Tailwind CSS-like utility class parsing. It enables components to support
CSS-like styling patterns while maintaining Go's type safety and performance.

# Features

• CSS-like utility class parsing (Tailwind-inspired)
• Class name concatenation and management
• Padding, margin, border, and styling utilities
• Color parsing for common color names
• Border radius and opacity parsing
• Component variant management

# Quick Start

Parse utility classes:

	styles := utils.ParseClasses("p-4 bg-blue rounded-lg")
	// styles.Padding will be layout.UniformInset(unit.Dp(16))
	// styles.Background will be blue color
	// styles.Radius will be unit.Dp(8)

Combine class names:

	className := utils.ClassNames("btn", "btn-primary", conditionalClass)

# Utility Classes

Supported utility class patterns:

Padding:
• p-{size} - Uniform padding
• px-{size}, py-{size} - Horizontal/vertical padding
• pt-{size}, pr-{size}, pb-{size}, pl-{size} - Individual sides

Margin:
• m-{size} - Uniform margin
• mx-{size}, my-{size} - Horizontal/vertical margin

Background:
• bg-{color} - Background color (red, blue, green, etc.)

Border:
• border - Default 1dp border
• border-{color} - Border color

Border Radius:
• rounded - Default 4dp radius
• rounded-{size} - Specific radius (sm, md, lg, xl, 2xl, 3xl, full)

Opacity:
• opacity-{value} - Opacity from 0-100

# Supported Sizes

Size scale (0-64):
• 0 = 0dp
• 1 = 4dp
• 2 = 8dp
• 3 = 12dp
• 4 = 16dp
• 6 = 24dp
• 8 = 32dp
• 12 = 48dp
• 16 = 64dp
• etc.

# Examples

Basic padding and background:

	styles := utils.ParseClasses("p-6 bg-blue rounded")
	// Applies 24dp padding, blue background, 4dp radius

Complex styling:

	styles := utils.ParseClasses("px-4 py-2 bg-white border border-gray rounded-lg opacity-90")
	// Horizontal padding: 16dp
	// Vertical padding: 8dp
	// White background
	// Gray border
	// Large border radius
	// 90% opacity

Component integration:

	// Parse classes and apply to component
	styles := utils.ParseClasses(component.Classes)
	if styles.Padding != (layout.Inset{}) {
		// Apply custom padding
		padding = styles.Padding
	}
	if styles.Background.A > 0 {
		// Apply custom background color
		bgColor = styles.Background
	}
*/
package utils

import (
	"image/color"
	"strings"

	"gioui.org/layout"
	"gioui.org/unit"
)

// ClassNames merges class names, similar to clsx in JavaScript.
func ClassNames(classes ...string) string {
	var result []string
	for _, class := range classes {
		if class != "" {
			result = append(result, class)
		}
	}
	return strings.Join(result, " ")
}

// StyleUtility represents parsed utility classes.
type StyleUtility struct {
	Padding    layout.Inset
	Margin     layout.Inset
	Background color.NRGBA
	Border     BorderStyle
	Radius     unit.Dp
	Width      unit.Dp
	Height     unit.Dp
	Opacity    float32
}

// BorderStyle represents border styling.
type BorderStyle struct {
	Color color.NRGBA
	Width unit.Dp
}

// ParseClasses parses Tailwind-like utility classes.
func ParseClasses(classes ...string) StyleUtility {
	style := StyleUtility{
		Opacity: 1.0,
	}

	classString := ClassNames(classes...)
	classList := strings.Fields(classString)

	for _, class := range classList {
		parseClass(class, &style)
	}

	return style
}

//nolint:gocyclo // This function has high complexity but is straightforward switch-based parsing
func parseClass(class string, style *StyleUtility) {
	switch {
	// Padding classes
	case strings.HasPrefix(class, "p-"):
		if spacing := parseSpacing(class[2:]); spacing != nil {
			style.Padding = layout.UniformInset(*spacing)
		}
	case strings.HasPrefix(class, "px-"):
		if spacing := parseSpacing(class[3:]); spacing != nil {
			style.Padding.Left = *spacing
			style.Padding.Right = *spacing
		}
	case strings.HasPrefix(class, "py-"):
		if spacing := parseSpacing(class[3:]); spacing != nil {
			style.Padding.Top = *spacing
			style.Padding.Bottom = *spacing
		}
	case strings.HasPrefix(class, "pt-"):
		if spacing := parseSpacing(class[3:]); spacing != nil {
			style.Padding.Top = *spacing
		}
	case strings.HasPrefix(class, "pr-"):
		if spacing := parseSpacing(class[3:]); spacing != nil {
			style.Padding.Right = *spacing
		}
	case strings.HasPrefix(class, "pb-"):
		if spacing := parseSpacing(class[3:]); spacing != nil {
			style.Padding.Bottom = *spacing
		}
	case strings.HasPrefix(class, "pl-"):
		if spacing := parseSpacing(class[3:]); spacing != nil {
			style.Padding.Left = *spacing
		}

	// Margin classes
	case strings.HasPrefix(class, "m-"):
		if spacing := parseSpacing(class[2:]); spacing != nil {
			style.Margin = layout.UniformInset(*spacing)
		}
	case strings.HasPrefix(class, "mx-"):
		if spacing := parseSpacing(class[3:]); spacing != nil {
			style.Margin.Left = *spacing
			style.Margin.Right = *spacing
		}
	case strings.HasPrefix(class, "my-"):
		if spacing := parseSpacing(class[3:]); spacing != nil {
			style.Margin.Top = *spacing
			style.Margin.Bottom = *spacing
		}

	// Border radius classes
	case strings.HasPrefix(class, "rounded-"):
		if radius := parseRadius(class[8:]); radius != nil {
			style.Radius = *radius
		}
	case class == "rounded":
		style.Radius = unit.Dp(4)

	// Background classes
	case strings.HasPrefix(class, "bg-"):
		if bgColor := parseColor(class[3:]); bgColor != nil {
			style.Background = *bgColor
		}

	// Border classes
	case strings.HasPrefix(class, "border-"):
		if borderColor := parseColor(class[7:]); borderColor != nil {
			style.Border.Color = *borderColor
		}
	case class == "border":
		style.Border.Width = unit.Dp(1)

	// Opacity classes
	case strings.HasPrefix(class, "opacity-"):
		if opacity := parseOpacity(class[8:]); opacity != nil {
			style.Opacity = *opacity
		}
	}
}

func parseSpacing(value string) *unit.Dp {
	spacingMap := map[string]unit.Dp{
		"0":  unit.Dp(0),
		"1":  unit.Dp(4),
		"2":  unit.Dp(8),
		"3":  unit.Dp(12),
		"4":  unit.Dp(16),
		"5":  unit.Dp(20),
		"6":  unit.Dp(24),
		"8":  unit.Dp(32),
		"10": unit.Dp(40),
		"12": unit.Dp(48),
		"16": unit.Dp(64),
		"20": unit.Dp(80),
		"24": unit.Dp(96),
		"32": unit.Dp(128),
		"40": unit.Dp(160),
		"48": unit.Dp(192),
		"56": unit.Dp(224),
		"64": unit.Dp(256),
	}

	if dp, exists := spacingMap[value]; exists {
		return &dp
	}
	return nil
}

func parseRadius(value string) *unit.Dp {
	radiusMap := map[string]unit.Dp{
		"none": unit.Dp(0),
		"sm":   unit.Dp(2),
		"":     unit.Dp(4), // default rounded
		"md":   unit.Dp(6),
		"lg":   unit.Dp(8),
		"xl":   unit.Dp(12),
		"2xl":  unit.Dp(16),
		"3xl":  unit.Dp(24),
		"full": unit.Dp(9999),
	}

	if dp, exists := radiusMap[value]; exists {
		return &dp
	}
	return nil
}

func parseColor(value string) *color.NRGBA {
	colorMap := map[string]color.NRGBA{
		"transparent": {R: 0, G: 0, B: 0, A: 0},
		"black":       {R: 0, G: 0, B: 0, A: 255},
		"white":       {R: 255, G: 255, B: 255, A: 255},
		"red":         {R: 239, G: 68, B: 68, A: 255},
		"green":       {R: 34, G: 197, B: 94, A: 255},
		"blue":        {R: 59, G: 130, B: 246, A: 255},
		"yellow":      {R: 251, G: 191, B: 36, A: 255},
		"purple":      {R: 147, G: 51, B: 234, A: 255},
		"pink":        {R: 236, G: 72, B: 153, A: 255},
		"indigo":      {R: 99, G: 102, B: 241, A: 255},
		"gray":        {R: 156, G: 163, B: 175, A: 255},
	}

	if clr, exists := colorMap[value]; exists {
		return &clr
	}
	return nil
}

func parseOpacity(value string) *float32 {
	opacityMap := map[string]float32{
		"0":   0.0,
		"5":   0.05,
		"10":  0.1,
		"20":  0.2,
		"25":  0.25,
		"30":  0.3,
		"40":  0.4,
		"50":  0.5,
		"60":  0.6,
		"70":  0.7,
		"75":  0.75,
		"80":  0.8,
		"90":  0.9,
		"95":  0.95,
		"100": 1.0,
	}

	if opacity, exists := opacityMap[value]; exists {
		return &opacity
	}
	return nil
}
