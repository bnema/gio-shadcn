package label

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
	"image/color"
)

// Label represents a shadcn/ui label component
type Label struct {
	// Configuration
	Text      string
	TextStyle theme.TextStyle
	Classes   string
	Variant   theme.Variant
	Size      theme.Size
}

// LabelOption is a functional option for configuring Label components
type LabelOption func(*Label)

// WithLabelText sets the label text
func WithLabelText(text string) LabelOption {
	return func(l *Label) {
		l.Text = text
	}
}

// WithTextStyle sets the label text style
func WithTextStyle(style theme.TextStyle) LabelOption {
	return func(l *Label) {
		l.TextStyle = style
	}
}

// WithLabelClasses sets additional CSS-like classes
func WithLabelClasses(classes string) LabelOption {
	return func(l *Label) {
		l.Classes = classes
	}
}

// WithLabelVariant sets the label variant
func WithLabelVariant(variant theme.Variant) LabelOption {
	return func(l *Label) {
		l.Variant = variant
	}
}

// WithLabelSize sets the label size
func WithLabelSize(size theme.Size) LabelOption {
	return func(l *Label) {
		l.Size = size
	}
}

// NewLabel creates a new Label with the given options
func NewLabel(options ...LabelOption) *Label {
	l := &Label{
		Variant: theme.VariantDefault,
		Size:    theme.SizeDefault,
	}

	for _, option := range options {
		option(l)
	}

	return l
}

// Layout renders the label
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

// Update returns the component state (Label has no interactive state)
func (l *Label) Update(gtx layout.Context) theme.ComponentState {
	return &LabelState{
		active:   false,
		hovered:  false,
		pressed:  false,
		disabled: false,
	}
}

// LabelState implements ComponentState for Label
type LabelState struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

func (ls *LabelState) IsActive() bool {
	return ls.active
}

func (ls *LabelState) IsHovered() bool {
	return ls.hovered
}

func (ls *LabelState) IsPressed() bool {
	return ls.pressed
}

func (ls *LabelState) IsDisabled() bool {
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

// SetText sets the label text
func (l *Label) SetText(text string) {
	l.Text = text
}

// SetTextStyle sets the label text style
func (l *Label) SetTextStyle(style theme.TextStyle) {
	l.TextStyle = style
}

// Typography component for various text elements
type Typography struct {
	Text      string
	Element   TypographyElement
	Classes   string
	TextStyle theme.TextStyle
}

// TypographyElement represents different typography elements
type TypographyElement string

const (
	H1    TypographyElement = "h1"
	H2    TypographyElement = "h2"
	H3    TypographyElement = "h3"
	H4    TypographyElement = "h4"
	P     TypographyElement = "p"
	Small TypographyElement = "small"
	Lead  TypographyElement = "lead"
	Large TypographyElement = "large"
	Muted TypographyElement = "muted"
)

// NewTypography creates a new typography component
func NewTypography(text string, element TypographyElement, classes string) *Typography {
	return &Typography{
		Text:    text,
		Element: element,
		Classes: classes,
	}
}

// Layout renders the typography component
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

// SetText sets the typography text
func (t *Typography) SetText(text string) {
	t.Text = text
}

// SetElement sets the typography element
func (t *Typography) SetElement(element TypographyElement) {
	t.Element = element
}

// SetTextStyle sets the typography text style
func (t *Typography) SetTextStyle(style theme.TextStyle) {
	t.TextStyle = style
}
