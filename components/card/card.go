package card

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/bnema/gio-shadcn/theme"
	"github.com/bnema/gio-shadcn/utils"
	"image"
)

// Card represents a shadcn/ui card component
type Card struct {
	// Configuration
	Variant theme.Variant
	Classes string
	Padding layout.Inset
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
					Width: variant.BorderWidth,
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
	// This is a simplified text rendering - in a real implementation,
	// you'd want to use material.Label or a proper text renderer
	return layout.Dimensions{
		Size: gtx.Constraints.Min,
	}
}
