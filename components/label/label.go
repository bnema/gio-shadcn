/*
Package label provides typography and text labeling components for gio-shadcn applications.

The label component provides semantic text rendering with consistent typography,
variants, sizes, and theming. It supports different typography elements like
headings, body text, captions, and muted text following shadcn/ui design principles.

# Quick Start

Create a basic label:

	label := label.NewTypography("Hello World", label.P, "")

Create a heading:

	heading := label.NewTypography("Page Title", label.H1, "")

Create a label with custom styling:

	label := label.NewLabel(
		label.WithLabelText("Custom Label"),
		label.WithLabelVariant(theme.VariantSecondary),
		label.WithLabelSize(theme.SizeLG),
	)

# Typography Elements

Available typography elements:
• H1 - Main page headings (largest, bold)
• H2 - Section headings (large, semi-bold)
• H3 - Subsection headings (medium, semi-bold)
• H4 - Minor headings (base size, semi-bold)
• P - Body text (base size, normal weight)
• Small - Small text for captions and fine print
• Muted - Muted text for secondary information

# Variants

Text variants:
• VariantDefault - Standard text using foreground color
• VariantSecondary - Secondary text with muted color
• VariantDestructive - Error or warning text in red

# Sizes

Available sizes:
• SizeDefault - Standard text size
• SizeSM - Small text size
• SizeLG - Large text size

# Features

• Semantic typography elements (H1-H4, P, Small, Muted)
• Consistent text styling following shadcn/ui patterns
• Theme integration with automatic color adaptation
• Custom text styling support
• Variant-based color schemes
• Size-based font scaling
• CSS-like class utilities support

# Examples

Page title with H1:

	title := label.NewTypography("Welcome", label.H1, "")
	dims := title.Layout(gtx, th)

Body text paragraph:

	text := label.NewTypography("This is body text content.", label.P, "")
	dims := text.Layout(gtx, th)

Small caption text:

	caption := label.NewTypography("Image caption", label.Small, "")
	dims := caption.Layout(gtx, th)

Custom styled label:

	label := label.NewLabel(
		label.WithLabelText("Status: Active"),
		label.WithLabelVariant(theme.VariantDefault),
		label.WithTextStyle(theme.TextStyle{
			Size:   th.Typography.FontSizeLG,
			Weight: font.Bold,
			Color:  &th.Colors,
		}),
	)

Muted helper text:

	helper := label.NewTypography("This field is optional", label.Muted, "")
	dims := helper.Layout(gtx, th)
*/
package label

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Label represents a shadcn/ui label component.
type Label struct {
	// Configuration
	Text      string
	TextStyle theme.TextStyle
	Classes   string
	Variant   theme.Variant
	Size      theme.Size
}

// Option is a functional option for configuring Label components.
type Option func(*Label)

// WithLabelText sets the label text.
func WithLabelText(text string) Option {
	return func(l *Label) {
		l.Text = text
	}
}

// WithTextStyle sets the label text style.
func WithTextStyle(style theme.TextStyle) Option {
	return func(l *Label) {
		l.TextStyle = style
	}
}

// WithLabelClasses sets additional CSS-like classes.
func WithLabelClasses(classes string) Option {
	return func(l *Label) {
		l.Classes = classes
	}
}

// WithLabelVariant sets the label variant.
func WithLabelVariant(variant theme.Variant) Option {
	return func(l *Label) {
		l.Variant = variant
	}
}

// WithLabelSize sets the label size.
func WithLabelSize(size theme.Size) Option {
	return func(l *Label) {
		l.Size = size
	}
}

// NewLabel creates a new Label with the given options.
func NewLabel(options ...Option) *Label {
	l := &Label{
		Variant: theme.VariantDefault,
		Size:    theme.SizeDefault,
	}

	for _, option := range options {
		option(l)
	}

	return l
}

// Layout renders the label.
func (l *Label) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	// Parse additional classes
	styles := utils.ParseClasses(l.Classes)

	// Determine text style
	textStyle := l.TextStyle
	if textStyle == (theme.TextStyle{}) {
		textStyle = l.getDefaultTextStyle(th)
	}

	// Apply size if specified
	if l.Size != "" {
		textStyle = l.applySizeToTextStyle(textStyle, th)
	}

	// Create material label
	label := material.Label(material.NewTheme(), textStyle.Size, l.Text)

	// Apply text color
	if textStyle.Color != nil {
		label.Color = textStyle.Color.Foreground
	}

	// Apply alignment
	label.Alignment = textStyle.Alignment

	// Apply font weight and style
	label.Font.Weight = textStyle.Weight
	label.Font.Style = textStyle.Style

	// Apply custom color if specified in classes
	if styles.Background.A > 0 {
		// For labels, background in classes might represent text color
		label.Color = styles.Background
	}

	return label.Layout(gtx)
}

// Update returns the component state (Label has no interactive state).
func (l *Label) Update(_ layout.Context) theme.ComponentState {
	return &State{
		active:   false,
		hovered:  false,
		pressed:  false,
		disabled: false,
	}
}

// State implements ComponentState for Label.
type State struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

// IsActive returns true if the label is active (labels are never active).
func (ls *State) IsActive() bool {
	return ls.active
}

// IsHovered returns true if the label is being hovered over (labels are never hovered).
func (ls *State) IsHovered() bool {
	return ls.hovered
}

// IsPressed returns true if the label is being pressed (labels are never pressed).
func (ls *State) IsPressed() bool {
	return ls.pressed
}

// IsDisabled returns true if the label is disabled (labels are never disabled).
func (ls *State) IsDisabled() bool {
	return ls.disabled
}

func (l *Label) getDefaultTextStyle(th *theme.Theme) theme.TextStyle {
	switch l.Variant {
	case theme.VariantDefault:
		return th.Typography.Body(&th.Colors)
	case theme.VariantSecondary:
		return theme.TextStyle{
			Size:          th.Typography.FontSizeBase,
			Color:         &th.Colors,
			Weight:        th.Typography.Body(&th.Colors).Weight,
			Style:         th.Typography.Body(&th.Colors).Style,
			Alignment:     th.Typography.Body(&th.Colors).Alignment,
			LineHeight:    th.Typography.Body(&th.Colors).LineHeight,
			LetterSpacing: th.Typography.Body(&th.Colors).LetterSpacing,
		}
	default:
		return th.Typography.Body(&th.Colors)
	}
}

func (l *Label) applySizeToTextStyle(textStyle theme.TextStyle, th *theme.Theme) theme.TextStyle {
	switch l.Size {
	case theme.SizeSM:
		textStyle.Size = th.Typography.FontSizeSM
	case theme.SizeLG:
		textStyle.Size = th.Typography.FontSizeLG
	default:
		textStyle.Size = th.Typography.FontSizeBase
	}
	return textStyle
}

// SetText sets the label text.
func (l *Label) SetText(text string) {
	l.Text = text
}

// SetTextStyle sets the label text style.
func (l *Label) SetTextStyle(style theme.TextStyle) {
	l.TextStyle = style
}

// Typography component for various text elements.
type Typography struct {
	Text      string
	Element   TypographyElement
	Classes   string
	TextStyle theme.TextStyle
}

// TypographyElement represents different typography elements.
type TypographyElement string

const (
	// H1 represents a main heading typography element.
	H1 TypographyElement = "h1"
	// H2 represents a section heading typography element.
	H2 TypographyElement = "h2"
	// H3 represents a subsection heading typography element.
	H3 TypographyElement = "h3"
	// H4 represents a minor heading typography element.
	H4 TypographyElement = "h4"
	// P represents a paragraph typography element.
	P TypographyElement = "p"
	// Small represents small text typography element.
	Small TypographyElement = "small"
	// Lead represents lead text typography element.
	Lead TypographyElement = "lead"
	// Large represents large text typography element.
	Large TypographyElement = "large"
	// Muted represents muted text typography element.
	Muted TypographyElement = "muted"
)

// NewTypography creates a new typography component.
func NewTypography(text string, element TypographyElement, classes string) *Typography {
	return &Typography{
		Text:    text,
		Element: element,
		Classes: classes,
	}
}

// Layout renders the typography component.
func (t *Typography) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	// Parse additional classes
	styles := utils.ParseClasses(t.Classes)

	// Get text style based on element
	textStyle := t.getTextStyleForElement(th)

	// Override with custom text style if provided
	if t.TextStyle != (theme.TextStyle{}) {
		textStyle = t.TextStyle
	}

	// Create material label
	label := material.Label(material.NewTheme(), textStyle.Size, t.Text)

	// Apply text color
	if textStyle.Color != nil {
		label.Color = t.getColorForElement(th)
	}

	// Apply alignment
	label.Alignment = textStyle.Alignment

	// Apply font weight and style
	label.Font.Weight = textStyle.Weight
	label.Font.Style = textStyle.Style

	// Apply custom color if specified in classes
	if styles.Background.A > 0 {
		label.Color = styles.Background
	}

	return label.Layout(gtx)
}

func (t *Typography) getTextStyleForElement(th *theme.Theme) theme.TextStyle {
	switch t.Element {
	case H1:
		return th.Typography.H1(&th.Colors)
	case H2:
		return th.Typography.H2(&th.Colors)
	case H3:
		return th.Typography.H3(&th.Colors)
	case H4:
		return th.Typography.H4(&th.Colors)
	case P:
		return th.Typography.Body(&th.Colors)
	case Small:
		return th.Typography.BodySmall(&th.Colors)
	case Lead:
		style := th.Typography.Body(&th.Colors)
		style.Size = th.Typography.FontSizeLG
		return style
	case Large:
		style := th.Typography.Body(&th.Colors)
		style.Size = th.Typography.FontSizeXL
		return style
	case Muted:
		style := th.Typography.BodySmall(&th.Colors)
		return style
	default:
		return th.Typography.Body(&th.Colors)
	}
}

func (t *Typography) getColorForElement(th *theme.Theme) color.NRGBA {
	switch t.Element {
	case H1, H2, H3, H4:
		return th.Colors.Foreground
	case P, Small, Lead, Large:
		return th.Colors.Foreground
	case Muted:
		return th.Colors.MutedFg
	default:
		return th.Colors.Foreground
	}
}

// SetText sets the typography text.
func (t *Typography) SetText(text string) {
	t.Text = text
}

// SetElement sets the typography element.
func (t *Typography) SetElement(element TypographyElement) {
	t.Element = element
}

// SetTextStyle sets the typography text style.
func (t *Typography) SetTextStyle(style theme.TextStyle) {
	t.TextStyle = style
}
