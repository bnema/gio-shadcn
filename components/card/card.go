package card

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
)

// Card represents a shadcn/ui card component
type Card struct {
	// Configuration
	Variant theme.Variant
	Classes string
	Padding layout.Inset
}

// CardOption is a functional option for configuring Card components
type CardOption func(*Card)

// WithCardVariant sets the card variant
func WithCardVariant(variant theme.Variant) CardOption {
	return func(c *Card) {
		c.Variant = variant
	}
}

// WithCardClasses sets additional CSS-like classes
func WithCardClasses(classes string) CardOption {
	return func(c *Card) {
		c.Classes = classes
	}
}

// WithCardPadding sets custom padding
func WithCardPadding(padding layout.Inset) CardOption {
	return func(c *Card) {
		c.Padding = padding
	}
}

// NewCard creates a new Card with the given options
func NewCard(options ...CardOption) *Card {
	c := &Card{
		Variant: theme.VariantDefault,
		Padding: layout.Inset{Top: 24, Right: 24, Bottom: 24, Left: 24},
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// Config represents card configuration
type Config struct {
	Variant theme.Variant
	Classes string
	Padding layout.Inset
}

// New creates a new card with the given configuration
func New(config Config) *Card {
	return &Card{
		Variant: config.Variant,
		Classes: config.Classes,
		Padding: config.Padding,
	}
}

// Layout renders the card with the given content
func (c *Card) Layout(gtx layout.Context, th *theme.Theme, content layout.Widget) layout.Dimensions {
	// Get variant configuration
	variant := utils.GetCardVariant(c.Variant, &th.Colors)

	// Parse additional classes
	styles := utils.ParseClasses(c.Classes)

	// Determine padding
	padding := c.Padding
	if padding == (layout.Inset{}) {
		padding = layout.Inset{
			Top:    th.Spacing.Space6,
			Bottom: th.Spacing.Space6,
			Left:   th.Spacing.Space6,
			Right:  th.Spacing.Space6,
		}
	}

	// Apply custom padding if specified in classes
	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}

	// Determine background color
	bgColor := variant.Background
	if styles.Background.A > 0 {
		bgColor = styles.Background
	}

	// Determine border radius
	radius := th.Radius.RadiusLG
	if styles.Radius > 0 {
		radius = styles.Radius
	}

	return layout.Stack{}.Layout(gtx,
		// Background
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			dims := padding.Layout(gtx, content)

			// Draw background
			rect := image.Rectangle{Max: dims.Size}
			rr := clip.UniformRRect(rect, gtx.Dp(radius))
			paint.FillShape(gtx.Ops, bgColor, rr.Op(gtx.Ops))

			// Draw border
			if variant.BorderWidth > 0 {
				border := clip.Stroke{
					Path:  rr.Path(gtx.Ops),
					Width: float32(gtx.Dp(unit.Dp(variant.BorderWidth))),
				}
				paint.FillShape(gtx.Ops, variant.Border, border.Op())
			}

			return dims
		}),

		// Content
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, content)
		}),
	)
}

func (c *Card) Update(gtx layout.Context) theme.ComponentState {
	return &CardState{
		active:   false,
		hovered:  false,
		pressed:  false,
		disabled: false,
	}
}

type CardState struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

func (cs *CardState) IsActive() bool {
	return cs.active
}

func (cs *CardState) IsHovered() bool {
	return cs.hovered
}

func (cs *CardState) IsPressed() bool {
	return cs.pressed
}

func (cs *CardState) IsDisabled() bool {
	return cs.disabled
}

// CardHeader represents a card header component
type CardHeader struct {
	Classes string
	Padding layout.Inset
}

// NewHeader creates a new card header
func NewHeader(classes string) *CardHeader {
	return &CardHeader{
		Classes: classes,
	}
}

// Layout renders the card header
func (h *CardHeader) Layout(gtx layout.Context, th *theme.Theme, content layout.Widget) layout.Dimensions {
	// Parse additional classes
	styles := utils.ParseClasses(h.Classes)

	// Determine padding
	padding := h.Padding
	if padding == (layout.Inset{}) {
		padding = layout.Inset{
			Top:    th.Spacing.Space6,
			Bottom: th.Spacing.Space6,
			Left:   th.Spacing.Space6,
			Right:  th.Spacing.Space6,
		}
	}

	// Apply custom padding if specified in classes
	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}

	return padding.Layout(gtx, content)
}

// CardContent represents a card content component
type CardContent struct {
	Classes string
	Padding layout.Inset
}

// NewContent creates a new card content
func NewContent(classes string) *CardContent {
	return &CardContent{
		Classes: classes,
	}
}

// Layout renders the card content
func (c *CardContent) Layout(gtx layout.Context, th *theme.Theme, content layout.Widget) layout.Dimensions {
	// Parse additional classes
	styles := utils.ParseClasses(c.Classes)

	// Determine padding
	padding := c.Padding
	if padding == (layout.Inset{}) {
		padding = layout.Inset{
			Top:    th.Spacing.Space6,
			Bottom: th.Spacing.Space6,
			Left:   th.Spacing.Space6,
			Right:  th.Spacing.Space6,
		}
	}

	// Apply custom padding if specified in classes
	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}

	return padding.Layout(gtx, content)
}

// CardFooter represents a card footer component
type CardFooter struct {
	Classes string
	Padding layout.Inset
}

// NewFooter creates a new card footer
func NewFooter(classes string) *CardFooter {
	return &CardFooter{
		Classes: classes,
	}
}

// Layout renders the card footer
func (f *CardFooter) Layout(gtx layout.Context, th *theme.Theme, content layout.Widget) layout.Dimensions {
	// Parse additional classes
	styles := utils.ParseClasses(f.Classes)

	// Determine padding
	padding := f.Padding
	if padding == (layout.Inset{}) {
		padding = layout.Inset{
			Top:    th.Spacing.Space6,
			Bottom: th.Spacing.Space6,
			Left:   th.Spacing.Space6,
			Right:  th.Spacing.Space6,
		}
	}

	// Apply custom padding if specified in classes
	if styles.Padding != (layout.Inset{}) {
		padding = styles.Padding
	}

	return padding.Layout(gtx, content)
}

// CardTitle represents a card title component
type CardTitle struct {
	Text    string
	Classes string
}

// NewTitle creates a new card title
func NewTitle(text string, classes string) *CardTitle {
	return &CardTitle{
		Text:    text,
		Classes: classes,
	}
}

// Layout renders the card title
func (t *CardTitle) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	textStyle := th.Typography.H3(&th.Colors)

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Height: th.Spacing.Space1}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return renderText(gtx, textStyle, t.Text)
		}),
	)
}

// CardDescription represents a card description component
type CardDescription struct {
	Text    string
	Classes string
}

// NewDescription creates a new card description
func NewDescription(text string, classes string) *CardDescription {
	return &CardDescription{
		Text:    text,
		Classes: classes,
	}
}

// Layout renders the card description
func (d *CardDescription) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	textStyle := th.Typography.BodySmall(&th.Colors)

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Height: th.Spacing.Space2}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return renderText(gtx, textStyle, d.Text)
		}),
	)
}

// Helper function to render text
func renderText(gtx layout.Context, style theme.TextStyle, text string) layout.Dimensions {
	// Create a material theme and label for text rendering
	thMat := material.NewTheme()
	label := material.Label(thMat, style.Size, text)
	// Use foreground color from the color scheme
	if style.Color != nil {
		label.Color = style.Color.Foreground
	}

	// Apply font weight if available
	if style.Weight > 0 {
		label.Font.Weight = style.Weight
	}

	return label.Layout(gtx)
}
