package theme

import (
	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
)

// Typography defines the typography system
type Typography struct {
	// Font families
	FontSans  []font.Face
	FontMono  []font.Face
	FontSerif []font.Face

	// Font sizes
	FontSizeXS   unit.Sp // 12px
	FontSizeSM   unit.Sp // 14px
	FontSizeBase unit.Sp // 16px
	FontSizeLG   unit.Sp // 18px
	FontSizeXL   unit.Sp // 20px
	FontSize2XL  unit.Sp // 24px
	FontSize3XL  unit.Sp // 30px
	FontSize4XL  unit.Sp // 36px

	// Line heights
	LineHeightTight   float32
	LineHeightSnug    float32
	LineHeightNormal  float32
	LineHeightRelaxed float32
	LineHeightLoose   float32

	// Letter spacing
	LetterSpacingTighter float32
	LetterSpacingTight   float32
	LetterSpacingNormal  float32
	LetterSpacingWide    float32
	LetterSpacingWider   float32
	LetterSpacingWidest  float32
}

// TextStyle represents a complete text style configuration
type TextStyle struct {
	Size          unit.Sp
	Color         *ColorScheme
	Weight        font.Weight
	Style         font.Style
	Alignment     text.Alignment
	LineHeight    float32
	LetterSpacing float32
}

// DefaultTypography returns the default typography configuration
func DefaultTypography() Typography {
	return Typography{
		FontSans:  []font.Face{}, // Will use system default
		FontMono:  []font.Face{}, // Will use system default
		FontSerif: []font.Face{}, // Will use system default

		FontSizeXS:   unit.Sp(12),
		FontSizeSM:   unit.Sp(14),
		FontSizeBase: unit.Sp(16),
		FontSizeLG:   unit.Sp(18),
		FontSizeXL:   unit.Sp(20),
		FontSize2XL:  unit.Sp(24),
		FontSize3XL:  unit.Sp(30),
		FontSize4XL:  unit.Sp(36),

		LineHeightTight:   1.25,
		LineHeightSnug:    1.375,
		LineHeightNormal:  1.5,
		LineHeightRelaxed: 1.625,
		LineHeightLoose:   2.0,

		LetterSpacingTighter: -0.05,
		LetterSpacingTight:   -0.025,
		LetterSpacingNormal:  0,
		LetterSpacingWide:    0.025,
		LetterSpacingWider:   0.05,
		LetterSpacingWidest:  0.1,
	}
}

// Predefined text styles
func (t *Typography) H1(colors *ColorScheme) TextStyle {
	return TextStyle{
		Size:          t.FontSize4XL,
		Color:         colors,
		Weight:        font.Bold,
		Style:         font.Regular,
		Alignment:     text.Start,
		LineHeight:    t.LineHeightTight,
		LetterSpacing: t.LetterSpacingTighter,
	}
}

func (t *Typography) H2(colors *ColorScheme) TextStyle {
	return TextStyle{
		Size:          t.FontSize3XL,
		Color:         colors,
		Weight:        font.SemiBold,
		Style:         font.Regular,
		Alignment:     text.Start,
		LineHeight:    t.LineHeightTight,
		LetterSpacing: t.LetterSpacingTight,
	}
}

func (t *Typography) H3(colors *ColorScheme) TextStyle {
	return TextStyle{
		Size:          t.FontSize2XL,
		Color:         colors,
		Weight:        font.SemiBold,
		Style:         font.Regular,
		Alignment:     text.Start,
		LineHeight:    t.LineHeightTight,
		LetterSpacing: t.LetterSpacingNormal,
	}
}

func (t *Typography) H4(colors *ColorScheme) TextStyle {
	return TextStyle{
		Size:          t.FontSizeXL,
		Color:         colors,
		Weight:        font.SemiBold,
		Style:         font.Regular,
		Alignment:     text.Start,
		LineHeight:    t.LineHeightTight,
		LetterSpacing: t.LetterSpacingNormal,
	}
}

func (t *Typography) Body(colors *ColorScheme) TextStyle {
	return TextStyle{
		Size:          t.FontSizeBase,
		Color:         colors,
		Weight:        font.Normal,
		Style:         font.Regular,
		Alignment:     text.Start,
		LineHeight:    t.LineHeightNormal,
		LetterSpacing: t.LetterSpacingNormal,
	}
}

func (t *Typography) BodySmall(colors *ColorScheme) TextStyle {
	return TextStyle{
		Size:          t.FontSizeSM,
		Color:         colors,
		Weight:        font.Normal,
		Style:         font.Regular,
		Alignment:     text.Start,
		LineHeight:    t.LineHeightNormal,
		LetterSpacing: t.LetterSpacingNormal,
	}
}

func (t *Typography) Caption(colors *ColorScheme) TextStyle {
	return TextStyle{
		Size:          t.FontSizeXS,
		Color:         colors,
		Weight:        font.Normal,
		Style:         font.Regular,
		Alignment:     text.Start,
		LineHeight:    t.LineHeightNormal,
		LetterSpacing: t.LetterSpacingWide,
	}
}
