/*
Package input provides versatile text input components for gio-shadcn applications.

The input component supports multiple input types (text, password, number, email),
visual variants, sizes, and features like validation, placeholder text, and helper text.
It follows shadcn/ui design principles and integrates seamlessly with the theme system.

# Quick Start

Create a basic text input:

	input := input.Text("Enter your name...")

Create a password input:

	passwordInput := input.Password("Enter password")

Create an email input with validation:

	emailInput := input.Email("Enter email")

# Input Types

Available input types:
• InputText - Standard text input
• InputPassword - Password input with masked text
• InputNumber - Numeric input with validation
• InputEmail - Email input with validation

# Variants

Visual variants:
• InputDefault - Standard input with border
• InputFilled - Filled background input
• InputGhost - Minimal input without visible border

# Sizes

Available sizes:
• InputSizeSmall - Compact input for tight layouts
• InputSizeMedium - Standard size for most use cases
• InputSizeLarge - Large input for emphasis

# Features

• Multiple input types with appropriate validation
• Placeholder text support
• Helper text and labels
• Error state management
• Theme integration with automatic color adaptation
• Keyboard event handling
• Focus state management
• Change and submit callbacks

# Examples

Text input with label and helper:

	input := input.New(input.Config{
		Type: input.InputText,
		Placeholder: "Enter your name",
		Label: "Full Name",
		Helper: "Enter your first and last name",
	})

Email input with validation:

	emailInput := input.New(input.Config{
		Type: input.InputEmail,
		Placeholder: "user@example.com",
		Label: "Email Address",
		Required: true,
		OnChange: func(value string) {
			// Validate email format
		},
	})

Password input:

	passwordInput := input.New(input.Config{
		Type: input.InputPassword,
		Placeholder: "Enter secure password",
		Label: "Password",
		Helper: "Must be at least 8 characters",
		Required: true,
	})
*/
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

// Type represents the type of input field.
type Type string

// Input types.
const (
	InputText     Type = "text"
	InputPassword Type = "password"
	InputNumber   Type = "number"
	InputEmail    Type = "email"
)

// Variant represents the visual variant of the input.
type Variant string

// Input variants.
const (
	InputDefault Variant = "default"
	InputFilled  Variant = "filled"
	InputGhost   Variant = "ghost"
)

// Size represents the size of the input.
type Size string

// Input sizes.
const (
	InputSizeSmall  Size = "sm"
	InputSizeMedium Size = "md"
	InputSizeLarge  Size = "lg"
)

// Input represents a shadcn/ui input component.
type Input struct {
	// State
	editor widget.Editor

	// Configuration
	Placeholder string
	Value       string
	Type        Type
	Variant     Variant
	Size        Size
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

// Option is a functional option for configuring Input components.
type Option func(*Input)

// WithPlaceholder sets the input placeholder.
func WithPlaceholder(placeholder string) Option {
	return func(i *Input) {
		i.Placeholder = placeholder
	}
}

// WithInputType sets the input type.
func WithInputType(inputType Type) Option {
	return func(i *Input) {
		i.Type = inputType
	}
}

// WithInputVariant sets the input variant.
func WithInputVariant(variant Variant) Option {
	return func(i *Input) {
		i.Variant = variant
	}
}

// WithInputSize sets the input size.
func WithInputSize(size Size) Option {
	return func(i *Input) {
		i.Size = size
	}
}

// WithLabel sets the input label.
func WithLabel(label string) Option {
	return func(i *Input) {
		i.Label = label
	}
}

// WithHelper sets the helper text.
func WithHelper(helper string) Option {
	return func(i *Input) {
		i.Helper = helper
	}
}

// WithRequired sets the required state.
func WithRequired(required bool) Option {
	return func(i *Input) {
		i.Required = required
	}
}

// WithInputDisabled sets the disabled state.
func WithInputDisabled(disabled bool) Option {
	return func(i *Input) {
		i.Disabled = disabled
	}
}

// WithOnChange sets the change callback.
func WithOnChange(onChange func(string)) Option {
	return func(i *Input) {
		i.OnChange = onChange
	}
}

// WithOnSubmit sets the submit callback.
func WithOnSubmit(onSubmit func()) Option {
	return func(i *Input) {
		i.OnSubmit = onSubmit
	}
}

// NewInput creates a new Input with the given options.
func NewInput(options ...Option) *Input {
	i := &Input{
		Type:    InputText,
		Variant: InputDefault,
		Size:    InputSizeMedium,
		editor:  widget.Editor{},
	}

	for _, option := range options {
		option(i)
	}

	return i
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

// Update returns the component state for Input.
func (i *Input) Update(_ layout.Context) theme.ComponentState {
	return &State{
		active:   i.focused,
		hovered:  false, // TODO: Add hover detection
		pressed:  false,
		disabled: i.Disabled,
	}
}

// State implements ComponentState for Input.
type State struct {
	active   bool
	hovered  bool
	pressed  bool
	disabled bool
}

// IsActive returns true if the input is active (focused).
func (is *State) IsActive() bool {
	return is.active
}

// IsHovered returns true if the input is being hovered over.
func (is *State) IsHovered() bool {
	return is.hovered
}

// IsPressed returns true if the input is being pressed.
func (is *State) IsPressed() bool {
	return is.pressed
}

// IsDisabled returns true if the input is disabled.
func (is *State) IsDisabled() bool {
	return is.disabled
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
	return NewInput(
		WithPlaceholder(placeholder),
		WithInputType(InputText),
	)
}

// Password creates a password input with the given placeholder.
func Password(placeholder string) *Input {
	return NewInput(
		WithPlaceholder(placeholder),
		WithInputType(InputPassword),
	)
}

// Number creates a number input with the given placeholder.
func Number(placeholder string) *Input {
	return NewInput(
		WithPlaceholder(placeholder),
		WithInputType(InputNumber),
	)
}

// Email creates an email input with the given placeholder.
func Email(placeholder string) *Input {
	return NewInput(
		WithPlaceholder(placeholder),
		WithInputType(InputEmail),
	)
}

// WithVariant sets the input variant.
func (i *Input) WithVariant(variant Variant) *Input {
	i.Variant = variant
	return i
}

// WithSize sets the input size.
func (i *Input) WithSize(size Size) *Input {
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
