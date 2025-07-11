// Package input provides shadcn/ui input components for Gio.
package input

import (
	"image"
	"image/color"

	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/bnema/gio-shadcn/theme"
)

// InputType represents the type of input field.
type InputType string

// Input types.
const (
	InputText     InputType = "text"
	InputPassword InputType = "password"
	InputNumber   InputType = "number"
	InputEmail    InputType = "email"
)

// InputVariant represents the visual variant of the input.
type InputVariant string

// Input variants.
const (
	InputDefault InputVariant = "default"
	InputFilled  InputVariant = "filled"
	InputGhost   InputVariant = "ghost"
)

// InputSize represents the size of the input.
type InputSize string

// Input sizes.
const (
	InputSizeSmall  InputSize = "sm"
	InputSizeMedium InputSize = "md"
	InputSizeLarge  InputSize = "lg"
)

// Input represents a shadcn/ui input component.
type Input struct {
	// State
	editor widget.Editor

	// Configuration
	Placeholder string
	Value       string
	Type        InputType
	Variant     InputVariant
	Size        InputSize
	Disabled    bool
	Error       bool
	Required    bool
	Label       string
	Helper      string
	ErrorMsg    string

	// Callbacks
	OnChange func(string)
	OnFocus  func()
	OnBlur   func()
	OnSubmit func()

	// Internal
	lastValue string
	focused   bool
}

// New creates a new input component.
func New() *Input {
	return &Input{
		Type:    InputText,
		Variant: InputDefault,
		Size:    InputSizeMedium,
		editor:  widget.Editor{},
	}
}

// SetText sets the text content of the input.
func (i *Input) SetText(text string) {
	i.editor.SetText(text)
	i.Value = text
	i.lastValue = text
}

// Text returns the current text content of the input.
func (i *Input) Text() string {
	return i.editor.Text()
}

// Layout renders the input component.
func (i *Input) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	// Configure editor based on type
	i.configureEditor()

	// Process editor events (this handles all keyboard input automatically)
	for {
		event, ok := i.editor.Update(gtx)
		if !ok {
			break
		}
		switch event.(type) {
		case widget.SubmitEvent:
			if i.OnSubmit != nil {
				i.OnSubmit()
			}
		case widget.ChangeEvent:
			// Handle change events if needed
		}
	}

	// Handle focus events separately for UI state tracking
	for {
		event, ok := gtx.Event(key.FocusFilter{Target: &i.editor})
		if !ok {
			break
		}
		if e, ok := event.(key.FocusEvent); ok {
			i.focused = e.Focus
			if e.Focus && i.OnFocus != nil {
				i.OnFocus()
			} else if !e.Focus && i.OnBlur != nil {
				i.OnBlur()
			}
		}
	}

	// Check for text changes
	currentText := i.editor.Text()
	if currentText != i.lastValue {
		i.lastValue = currentText
		i.Value = currentText
		if i.OnChange != nil {
			i.OnChange(currentText)
		}
	}

	// Create editor style
	thMat := material.NewTheme()
	editor := material.Editor(thMat, &i.editor, i.Placeholder)
	editor.Color = i.getTextColor(th)
	editor.HintColor = th.Colors.MutedFg
	editor.TextSize = unit.Sp(14)

	// Calculate input dimensions based on size
	inputHeight := i.getInputHeight()
	padding := unit.Dp(12)

	// Set minimum height for the context
	minHeight := gtx.Dp(inputHeight)
	gtx.Constraints.Min.Y = minHeight

	// Calculate the bounds for background/border
	bounds := image.Rectangle{Max: image.Point{X: gtx.Constraints.Max.X, Y: minHeight}}

	// Draw background FIRST (behind the text)
	paint.FillShape(gtx.Ops, i.getBackgroundColor(th),
		clip.UniformRRect(bounds, gtx.Metric.Dp(6)).Op(gtx.Ops))

	// Use thicker border when focused
	borderWidth := unit.Dp(1)
	if i.focused {
		borderWidth = unit.Dp(2)
	}

	// Draw border SECOND (behind the text)
	paint.FillShape(gtx.Ops, i.getBorderColor(th),
		clip.Stroke{
			Path:  clip.UniformRRect(bounds, gtx.Metric.Dp(6)).Path(gtx.Ops),
			Width: float32(gtx.Metric.Dp(borderWidth)),
		}.Op())

	// Layout the editor with padding LAST (in front of background)
	dims := layout.UniformInset(padding).Layout(gtx, editor.Layout)

	// Ensure the final dimensions match our minimum height
	if dims.Size.Y < minHeight {
		dims.Size.Y = minHeight
	}

	return dims
}

func (i *Input) configureEditor() {
	switch i.Type {
	case InputPassword:
		i.editor.Mask = '*'
	case InputNumber:
		i.editor.Filter = "0123456789.-"
	case InputEmail:
		i.editor.Filter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@.-_"
	default:
		i.editor.Mask = 0
		i.editor.Filter = ""
	}

	i.editor.SingleLine = true
	i.editor.ReadOnly = i.Disabled
}

func (i *Input) getBackgroundColor(th *theme.Theme) color.NRGBA {
	if i.Disabled {
		return th.Colors.Muted
	}

	switch i.Variant {
	case InputFilled:
		return th.Colors.Secondary
	case InputGhost:
		return color.NRGBA{0, 0, 0, 0} // Transparent
	default:
		return th.Colors.Background
	}
}

func (i *Input) getBorderColor(th *theme.Theme) color.NRGBA {
	if i.Error {
		return th.Colors.Destructive
	}

	if i.Disabled {
		return th.Colors.Border
	}

	// Check if input is focused
	if i.focused {
		return th.Colors.Ring // Use ring color for focus state
	}

	return th.Colors.Border
}

func (i *Input) getTextColor(th *theme.Theme) color.NRGBA {
	if i.Disabled {
		return th.Colors.MutedFg
	}
	return th.Colors.Foreground
}

func (i *Input) getInputHeight() unit.Dp {
	switch i.Size {
	case InputSizeSmall:
		return unit.Dp(36)
	case InputSizeLarge:
		return unit.Dp(52)
	default: // InputSizeMedium
		return unit.Dp(44)
	}
}

// Text creates a text input with the given placeholder.
func Text(placeholder string) *Input {
	i := New()
	i.Placeholder = placeholder
	i.Type = InputText
	return i
}

// Password creates a password input with the given placeholder.
func Password(placeholder string) *Input {
	i := New()
	i.Placeholder = placeholder
	i.Type = InputPassword
	return i
}

// Number creates a number input with the given placeholder.
func Number(placeholder string) *Input {
	i := New()
	i.Placeholder = placeholder
	i.Type = InputNumber
	return i
}

// Email creates an email input with the given placeholder.
func Email(placeholder string) *Input {
	i := New()
	i.Placeholder = placeholder
	i.Type = InputEmail
	return i
}

// WithVariant sets the input variant.
func (i *Input) WithVariant(variant InputVariant) *Input {
	i.Variant = variant
	return i
}

// WithSize sets the input size.
func (i *Input) WithSize(size InputSize) *Input {
	i.Size = size
	return i
}

// WithLabel sets the input label.
func (i *Input) WithLabel(label string) *Input {
	i.Label = label
	return i
}

// WithHelper sets the input helper text.
func (i *Input) WithHelper(helper string) *Input {
	i.Helper = helper
	return i
}

// WithError sets the input error state and message.
func (i *Input) WithError(errorMsg string) *Input {
	i.Error = true
	i.ErrorMsg = errorMsg
	return i
}

// WithRequired sets the input required state.
func (i *Input) WithRequired(required bool) *Input {
	i.Required = required
	return i
}

// WithDisabled sets the input disabled state.
func (i *Input) WithDisabled(disabled bool) *Input {
	i.Disabled = disabled
	return i
}

// WithOnChange sets the input change callback.
func (i *Input) WithOnChange(fn func(string)) *Input {
	i.OnChange = fn
	return i
}

// WithOnFocus sets the input focus callback.
func (i *Input) WithOnFocus(fn func()) *Input {
	i.OnFocus = fn
	return i
}

// WithOnBlur sets the input blur callback.
func (i *Input) WithOnBlur(fn func()) *Input {
	i.OnBlur = fn
	return i
}

// WithOnSubmit sets the input submit callback.
func (i *Input) WithOnSubmit(fn func()) *Input {
	i.OnSubmit = fn
	return i
}
