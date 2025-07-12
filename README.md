# gio-shadcn

A Go port of [shadcn/ui](https://ui.shadcn.com/) for [Gio](https://gioui.org/), bringing beautiful, accessible, and customizable UI components to your Gio applications.

![Demo App](assets/demo-app.png)

## Overview

gio-shadcn provides a collection of reusable, themeable UI components for Gio applications. Each component follows consistent design patterns, making it easy to build cohesive user interfaces with minimal effort.

### Key Benefits
- **🎨 Consistent Design**: All components follow shadcn/ui design principles
- **🔧 Highly Customizable**: Flexible theming system with runtime theme switching
- **📦 Modular Architecture**: Import only what you need
- **🚀 Developer Friendly**: Functional options pattern for intuitive API
- **💪 Type Safe**: Full Go type safety with compile-time checks

## Features

- 🎨 **Themeable**: JSON-based theme configuration with light/dark mode support
- 🧩 **Modular**: Each component is independently importable
- 💻 **Simple Imports**: Standard Go module imports, no CLI needed
- 🔧 **Type Safe**: Full Go type safety with validation
- 🎯 **Gio Native**: Built specifically for Gio's immediate-mode architecture

## Work in Progress

⚠️ **This project is under active development.** While functional, some features are still being implemented:
- Only 5 of the planned 51 components are currently available
- More components are being added regularly (help wanted!)
- API may change before v1.0 release

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Quick Start](#quick-start)
  - [Installation](#installation)
  - [Basic Usage](#basic-usage)
- [Component Guide](#component-guide)
  - [Understanding the Component API](#understanding-the-component-api)
  - [Component Lifecycle](#component-lifecycle)
  - [Available Components](#available-components)
    - [Button](#button)
    - [Card](#card)
    - [Input](#input)
    - [Label](#label)
    - [Titlebar](#titlebar)
- [Theming & Customization](#theming--customization)
  - [Understanding the Theme System](#understanding-the-theme-system)
  - [Creating Custom Themes](#creating-custom-themes)
  - [Component Styling](#component-styling)
  - [Dynamic Theme Switching](#dynamic-theme-switching)
- [Architecture](#architecture)
  - [Design Principles](#design-principles)
  - [Component Interface](#component-interface)
  - [State Management](#state-management)
  - [Project Structure](#project-structure)
- [Examples](#examples)
  - [Form Example](#form-example)
  - [Dashboard Layout](#dashboard-layout)
- [Component Progress](#component-progress)
  - [Ported Components](#ported-components-551)
  - [High Priority Components](#high-priority-components)
  - [Missing Components](#missing-components-4651)
- [Development](#development)
  - [Building from Source](#building-from-source)
  - [Creating a New Component](#creating-a-new-component)
  - [Code Style Guidelines](#code-style-guidelines)
- [Contributing](#contributing)
  - [Contribution Guidelines](#contribution-guidelines)
  - [Priority Areas](#priority-areas)
- [Troubleshooting](#troubleshooting)
- [License](#license)
- [Credits](#credits)
- [Support](#support)

## Quick Start

### Installation

Add gio-shadcn to your Go project:
```bash
go get github.com/bnema/gio-shadcn@latest
```

### Basic Usage

Here's a minimal example to get you started:

```go
package main

import (
    "os"
    
    "gioui.org/app"
    "gioui.org/io/system"
    "gioui.org/layout"
    "gioui.org/op"
    
    "github.com/bnema/gio-shadcn/components/button"
    "github.com/bnema/gio-shadcn/theme"
)

func main() {
    go func() {
        // Create window
        w := app.NewWindow(app.Title("gio-shadcn Example"))
        defer w.Close()
        
        // Create theme
        th := theme.New()
        
        // Create button with functional options
        btn := button.NewButton(
            button.WithText("Click me"),
            button.WithVariant(theme.VariantDefault),
            button.WithOnClick(func() {
                println("Button clicked!")
            }),
        )
        
        // Event loop
        var ops op.Ops
        for {
            e := w.NextEvent()
            switch e := e.(type) {
            case system.DestroyEvent:
                return
            case system.FrameEvent:
                gtx := layout.NewContext(&ops, e)
                
                // Layout the button centered
                layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
                    return btn.Layout(gtx, th)
                })
                
                e.Frame(gtx.Ops)
            }
        }
    }()
    app.Main()
}
```

## Component Guide

### Understanding the Component API

All gio-shadcn components follow a consistent API pattern:

1. **Creation**: Use `NewComponent()` with functional options
2. **Layout**: Call `Layout(gtx, theme)` to render
3. **State**: Use `Update(gtx)` to get component state (for interactive components)
4. **Customization**: Apply theme variants, sizes, and custom styles

### Component Lifecycle

```go
// 1. Create component with options
component := package.NewComponent(
    package.WithOption1(value1),
    package.WithOption2(value2),
)

// 2. In your layout function
dims := component.Layout(gtx, theme)

// 3. Check state (if interactive)
state := component.Update(gtx)
if state.IsPressed() {
    // Handle interaction
}
```

### Available Components

#### Button

Versatile button component with multiple variants and sizes.

```go
// Create button with all available options
btn := button.NewButton(
    button.WithText("Click me"),
    button.WithVariant(theme.VariantDefault), // Default, Destructive, Outline, Secondary, Ghost, Link
    button.WithSize(theme.SizeDefault),       // Default, SM, LG, Icon
    button.WithOnClick(func() { 
        println("Clicked!") 
    }),
    button.WithDisabled(false),
    button.WithClasses("custom-class"),
    button.WithIcon(myIcon), // Optional icon widget
)

// Check button state
state := btn.Update(gtx)
if state.IsHovered() {
    // Show tooltip
}
```

**Variants explained:**
- `Default`: Primary action button with solid background
- `Destructive`: For dangerous actions (red theme)
- `Outline`: Border only, transparent background
- `Secondary`: Less prominent than default
- `Ghost`: Minimal styling, appears on hover
- `Link`: Styled like a hyperlink

#### Card

Container component for grouping related content.

```go
// Create card with padding
card := card.NewCard(
    card.WithCardVariant(theme.VariantDefault),
    card.WithCardPadding(layout.Inset{
        Top: 24, Right: 24, Bottom: 24, Left: 24,
    }),
    card.WithCardClasses("custom-card"),
)

// Use card as a container
dims := card.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
    // Your content here
    return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
        layout.Rigid(headerWidget),
        layout.Rigid(contentWidget),
    )
})
```

#### Input

Text input component with validation and different types.

```go
// Create input with all options
input := input.NewInput(
    input.WithPlaceholder("Enter your email"),
    input.WithInputType(input.InputEmail),      // Text, Password, Number, Email
    input.WithInputVariant(input.InputDefault), // Default, Filled, Ghost
    input.WithInputSize(input.InputSizeMedium), // Small, Medium, Large
    input.WithLabel("Email Address"),
    input.WithHelper("We'll never share your email"),
    input.WithRequired(true),
    input.WithInputDisabled(false),
    input.WithOnChange(func(value string) {
        // Validate as user types
        println("Current value:", value)
    }),
    input.WithOnSubmit(func() {
        // Handle form submission
        println("Submitted!")
    }),
)

// Get current value
currentValue := input.Text()

// Set value programmatically
input.SetText("user@example.com")

// Convenience constructors for common types
passwordInput := input.Password("Enter password")
emailInput := input.Email("Enter email")
numberInput := input.Number("Enter age")
```

#### Label

Typography component for displaying text with consistent styling.

```go
// Create label with styling
label := label.NewLabel(
    label.WithLabelText("Welcome to gio-shadcn"),
    label.WithLabelVariant(theme.VariantDefault), // Default, Secondary
    label.WithLabelSize(theme.SizeLG),            // Default, SM, LG
    label.WithLabelClasses("custom-label"),
    label.WithTextStyle(theme.TextStyle{
        Size:   th.Typography.FontSizeXL,
        Weight: font.Bold,
        Color:  &th.Colors.Primary,
    }),
)

// Convenience constructors for common typography
h1 := label.NewTypography("Main Title", label.H1, "")
paragraph := label.NewTypography("Body text", label.P, "")
small := label.NewTypography("Fine print", label.Small, "")
```

#### Titlebar

Custom window titlebar component for frameless windows.

```go
// Create titlebar
titlebar := titlebar.NewTitleBar(
    titlebar.WithTitle("My Application"),
    titlebar.WithWindow(window), // Your *app.Window
    titlebar.WithCloseHandler(func() {
        // Custom close logic
        if unsavedChanges {
            showSaveDialog()
        } else {
            os.Exit(0)
        }
    }),
)

// In your layout (window parameter required for compatibility)
dims := titlebar.Layout(gtx, th, window)
```

## Theming & Customization

### Understanding the Theme System

The theme system in gio-shadcn provides comprehensive control over your application's appearance.

#### Theme Structure

```go
type Theme struct {
    Colors     ColorPalette  // Color definitions
    Typography Typography    // Font settings
    Spacing    Spacing       // Spacing scale
    Radius     Radius        // Border radius scale
    IsDark     bool          // Dark mode flag
}
```

### Creating Custom Themes

#### Method 1: Programmatic Theme

```go
// Create default theme and modify
th := theme.New()
th.Colors.Primary = color.NRGBA{R: 59, G: 130, B: 246, A: 255}
th.Typography.FontSizeLG = unit.Sp(20)
th.Radius.MD = unit.Dp(8)
```

#### Method 2: JSON Theme File

Create a `theme.json` file:

```json
{
  "colors": {
    "light": {
      "background": "#ffffff",
      "foreground": "#0a0a0a",
      "primary": "#3b82f6",
      "primary-foreground": "#ffffff",
      "secondary": "#e5e7eb",
      "secondary-foreground": "#0a0a0a",
      "destructive": "#ef4444",
      "destructive-foreground": "#ffffff",
      "muted": "#f3f4f6",
      "muted-foreground": "#6b7280",
      "accent": "#f3f4f6",
      "accent-foreground": "#0a0a0a",
      "border": "#e5e7eb",
      "input": "#e5e7eb",
      "ring": "#3b82f6"
    },
    "dark": {
      "background": "#0a0a0a",
      "foreground": "#fafafa",
      "primary": "#3b82f6",
      "primary-foreground": "#ffffff",
      "secondary": "#262626",
      "secondary-foreground": "#fafafa",
      "destructive": "#7f1d1d",
      "destructive-foreground": "#fafafa",
      "muted": "#262626",
      "muted-foreground": "#a1a1aa",
      "accent": "#262626",
      "accent-foreground": "#fafafa",
      "border": "#262626",
      "input": "#262626",
      "ring": "#3b82f6"
    }
  },
  "typography": {
    "fontFamily": "system-ui",
    "fontSize": {
      "xs": 12,
      "sm": 14,
      "base": 16,
      "lg": 18,
      "xl": 20,
      "2xl": 24,
      "3xl": 30,
      "4xl": 36
    }
  },
  "spacing": {
    "0": 0,
    "1": 4,
    "2": 8,
    "3": 12,
    "4": 16,
    "5": 20,
    "6": 24,
    "8": 32,
    "10": 40,
    "12": 48,
    "16": 64,
    "20": 80,
    "24": 96
  },
  "radius": {
    "none": 0,
    "sm": 2,
    "md": 4,
    "lg": 8,
    "xl": 12,
    "2xl": 16,
    "3xl": 24,
    "full": 9999
  }
}
```

Load the theme:

```go
// Load theme from JSON
th, err := theme.NewThemeFromJSON("theme.json")
if err != nil {
    log.Fatal(err)
}

// Toggle dark mode
th.ToggleDark()

// Validate theme completeness
if err := theme.ValidateTheme(th); err != nil {
    log.Printf("Theme validation warning: %v", err)
}
```

### Component Styling

#### Using Theme Variants

Components support different visual variants that work with your theme:

```go
// Primary button (uses theme.Colors.Primary)
primaryBtn := button.NewButton(
    button.WithText("Save"),
    button.WithVariant(theme.VariantDefault),
)

// Danger button (uses theme.Colors.Destructive)
deleteBtn := button.NewButton(
    button.WithText("Delete"),
    button.WithVariant(theme.VariantDestructive),
)
```

#### Custom Styling

Apply custom styles while maintaining theme consistency:

```go
// Custom styled button
customBtn := button.NewButton(
    button.WithText("Custom"),
    button.WithClasses("my-custom-button"), // For identification
)

// In your layout, apply custom styling
dims := customBtn.Layout(gtx, th)

// Access component state for conditional styling
state := customBtn.Update(gtx)
if state.IsHovered() {
    // Apply hover effects
}
```

### Dynamic Theme Switching

```go
// Theme toggle button
var isDarkMode bool
themeToggle := button.NewButton(
    button.WithText("🌙 Dark Mode"),
    button.WithOnClick(func() {
        th.ToggleDark()
        isDarkMode = !isDarkMode
        if isDarkMode {
            themeToggle.SetText("☀️ Light Mode")
        } else {
            themeToggle.SetText("🌙 Dark Mode")
        }
        window.Invalidate() // Force redraw
    }),
)
```

## Architecture

### Design Principles

1. **Functional Options Pattern**: All components use functional options for configuration
2. **Immutable State**: Components manage their own state internally
3. **Theme Integration**: Every component respects the current theme
4. **Composability**: Components can be easily combined and nested

### Component Interface

All components implement the `Component` interface:

```go
type Component interface {
    Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions
    Update(gtx layout.Context) ComponentState
}
```

### State Management

Interactive components provide state information:

```go
type ComponentState interface {
    IsActive() bool
    IsHovered() bool  
    IsPressed() bool
    IsDisabled() bool
}
```

### Project Structure

```
gio-shadcn/
├── components/          # Component implementations
│   ├── button/
│   ├── card/
│   ├── input/
│   ├── label/
│   └── titlebar/
├── theme/              # Theme system
├── utils/              # Shared utilities
└── examples/           # Example applications
```

## Examples

### Form Example

Complete form with validation:

```go
func createLoginForm(th *theme.Theme) layout.Widget {
    // Form inputs
    emailInput := input.NewInput(
        input.WithPlaceholder("email@example.com"),
        input.WithInputType(input.InputEmail),
        input.WithLabel("Email"),
        input.WithRequired(true),
    )
    
    passwordInput := input.NewInput(
        input.WithPlaceholder("Enter password"),
        input.WithInputType(input.InputPassword),
        input.WithLabel("Password"),
        input.WithRequired(true),
        input.WithHelper("Must be at least 8 characters"),
    )
    
    // Error message label
    errorLabel := label.NewLabel(
        label.WithLabelText(""),
        label.WithLabelVariant(theme.VariantDestructive),
    )
    
    // Submit button
    submitBtn := button.NewButton(
        button.WithText("Sign In"),
        button.WithVariant(theme.VariantDefault),
        button.WithSize(theme.SizeLG),
        button.WithOnClick(func() {
            // Validate inputs
            if emailInput.Text() == "" || passwordInput.Text() == "" {
                errorLabel.SetText("Please fill in all fields")
                return
            }
            
            if len(passwordInput.Text()) < 8 {
                errorLabel.SetText("Password too short")
                return
            }
            
            // Process login
            errorLabel.SetText("")
            println("Login:", emailInput.Text())
        }),
    )
    
    // Form card
    formCard := card.NewCard(
        card.WithCardPadding(layout.Inset{
            Top: 32, Right: 32, Bottom: 32, Left: 32,
        }),
    )
    
    return func(gtx layout.Context) layout.Dimensions {
        return formCard.Layout(gtx, th, func(gtx layout.Context) layout.Dimensions {
            return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
                // Title
                layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                    title := label.NewTypography("Sign In", label.H2, "")
                    return title.Layout(gtx, th)
                }),
                
                layout.Rigid(layout.Spacer{Height: th.Spacing.Space6}.Layout),
                
                // Email input
                layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                    return emailInput.Layout(gtx, th)
                }),
                
                layout.Rigid(layout.Spacer{Height: th.Spacing.Space4}.Layout),
                
                // Password input
                layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                    return passwordInput.Layout(gtx, th)
                }),
                
                layout.Rigid(layout.Spacer{Height: th.Spacing.Space2}.Layout),
                
                // Error message
                layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                    return errorLabel.Layout(gtx, th)
                }),
                
                layout.Rigid(layout.Spacer{Height: th.Spacing.Space6}.Layout),
                
                // Submit button
                layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                    return submitBtn.Layout(gtx, th)
                }),
            )
        })
    }
}
```

### Dashboard Layout

```go
func createDashboard(th *theme.Theme) layout.Widget {
    // Stats cards
    statsCards := []struct {
        title string
        value string
        change string
    }{
        {"Total Revenue", "$45,231.89", "+20.1%"},
        {"Subscriptions", "+2,350", "+180.1%"},
        {"Sales", "+12,234", "+19.0%"},
        {"Active Now", "+573", "+201"},
    }
    
    return func(gtx layout.Context) layout.Dimensions {
        return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
            // Header
            layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                title := label.NewTypography("Dashboard", label.H1, "")
                subtitle := label.NewTypography("Welcome back! Here's your overview.", label.P, "")
                
                return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
                    layout.Rigid(title.Layout),
                    layout.Rigid(layout.Spacer{Height: th.Spacing.Space2}.Layout),
                    layout.Rigid(subtitle.Layout),
                )
            }),
            
            layout.Rigid(layout.Spacer{Height: th.Spacing.Space6}.Layout),
            
            // Stats grid
            layout.Rigid(func(gtx layout.Context) layout.Dimensions {
                return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
                    // Create stat cards
                    layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
                        // Grid layout for cards
                        // ... card layout code
                    }),
                )
            }),
        )
    }
}
```

## Component Progress

### ✅ Ported Components (5/51)

| Component | Import Path | Status | Description |
|-----------|-------------|--------|-------------|
| Button | `github.com/bnema/gio-shadcn/components/button` | ✅ Complete | Customizable button with variants and sizes |
| Card | `github.com/bnema/gio-shadcn/components/card` | ✅ Complete | Flexible container for content |
| Input | `github.com/bnema/gio-shadcn/components/input` | ✅ Complete | Text input with validation |
| Label | `github.com/bnema/gio-shadcn/components/label` | ✅ Complete | Typography component |
| Titlebar | `github.com/bnema/gio-shadcn/components/titlebar` | ✅ Complete | Window titlebar component |

### 🚧 High Priority Components

These components are most requested and will be implemented next:

| Component | Description | Use Case |
|-----------|-------------|----------|
| Select | Dropdown selection | Form inputs, settings |
| Checkbox | Boolean input | Forms, preferences |
| Switch | Toggle switch | Settings, feature flags |
| Dialog | Modal dialogs | Confirmations, forms |
| Tabs | Tabbed content | Navigation, organization |
| Tooltip | Hover information | Help text, hints |

See the full list of [planned components](#missing-components-4651) below.

### 🚧 Missing Components (46/51)

| Component | Priority | Description |
|-----------|----------|-------------|
| Accordion | High | Collapsible content areas |
| Alert | High | Display important messages |
| Alert Dialog | High | Modal dialog for alerts |
| Aspect Ratio | Medium | Maintain aspect ratios |
| Avatar | Medium | User profile pictures |
| Badge | High | Small status indicators |
| Breadcrumb | Medium | Navigation breadcrumbs |
| Calendar | Medium | Date selection |
| Carousel | Low | Image/content slider |
| Chart | Low | Data visualization |
| Checkbox | High | Boolean input control |
| Collapsible | Medium | Expandable content |
| Combobox | High | Searchable select |
| Command | Medium | Command palette |
| Context Menu | Medium | Right-click menus |
| Data Table | High | Tabular data display |
| Date Picker | High | Date selection input |
| Dialog | High | Modal dialogs |
| Drawer | Medium | Slide-out panels |
| Dropdown Menu | High | Dropdown selections |
| React Hook Form | N/A | Form handling (React specific) |
| Hover Card | Low | Hover tooltips |
| Input OTP | Medium | One-time password input |
| Menubar | Medium | Application menu bar |
| Navigation Menu | High | Site navigation |
| Pagination | High | Page navigation |
| Popover | Medium | Floating content |
| Progress | High | Progress indicators |
| Radio Group | High | Single selection from options |
| Resizable | Low | Resizable panels |
| Scroll-area | Medium | Custom scrollbars |
| Select | High | Dropdown selection |
| Separator | High | Visual dividers |
| Sheet | Medium | Side panels |
| Sidebar | High | Navigation sidebar |
| Skeleton | Medium | Loading placeholders |
| Slider | High | Range input control |
| Sonner | Low | Toast notifications |
| Switch | High | Toggle switch |
| Table | High | Data tables |
| Tabs | High | Tabbed content |
| Textarea | High | Multi-line text input |
| Toast | High | Notification messages |
| Toggle | High | Toggle button |
| Toggle Group | Medium | Group of toggle buttons |
| Tooltip | High | Hover information |
| Typography | Medium | Text styling utilities |

**Progress: 5/51 components ported (9.8%)**

## Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/bnema/gio-shadcn.git
cd gio-shadcn

# Install dependencies
go mod download

# Run the demo
go run ./cmd/demo-app

# Run tests
go test ./...
```

### Creating a New Component

Follow these steps to add a new component:

1. **Create component directory**:
```bash
mkdir components/newcomponent
cd components/newcomponent
```

2. **Implement the component**:
```go
// newcomponent.go
package newcomponent

import (
    "gioui.org/layout"
    "github.com/bnema/gio-shadcn/theme"
)

// NewComponent state
type NewComponent struct {
    // Add fields
}

// Functional options
type Option func(*NewComponent)

func WithText(text string) Option {
    return func(nc *NewComponent) {
        nc.text = text
    }
}

// Constructor
func NewNewComponent(opts ...Option) *NewComponent {
    nc := &NewComponent{
        // Default values
    }
    
    for _, opt := range opts {
        opt(nc)
    }
    
    return nc
}

// Layout implements Component interface
func (nc *NewComponent) Layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
    // Implementation
}

// Update implements Component interface
func (nc *NewComponent) Update(gtx layout.Context) theme.ComponentState {
    return &NewComponentState{
        // State values
    }
}
```

3. **Add tests**:
```go
// newcomponent_test.go
package newcomponent

import "testing"

func TestNewComponent(t *testing.T) {
    // Test implementation
}
```

4. **Update documentation**:
- Add to README component list
- Create example usage
- Update progress tracking

### Code Style Guidelines

- Follow standard Go conventions
- Use functional options for all configuration
- Implement the Component interface
- Include comprehensive documentation
- Add unit tests for new functionality

## Contributing

We welcome contributions! Here's how to get involved:

1. **Check existing issues** for components to implement
2. **Fork the repository** and create a feature branch
3. **Implement your component** following the patterns above
4. **Add tests and examples**
5. **Update documentation**
6. **Submit a pull request**

### Contribution Guidelines

- Maintain consistency with existing component APIs
- Follow the functional options pattern
- Ensure all components work with the theme system
- Include examples in your PR
- Test with both light and dark themes

### Priority Areas

- High-priority missing components (see list above)
- Theme system enhancements
- Performance optimizations
- Documentation improvements
- Example applications

## Troubleshooting

### Common Issues

**Q: Components not rendering correctly?**
- Ensure you're passing the theme to Layout()
- Check that your layout constraints are properly set
- Verify the component is receiving events (for interactive components)

**Q: Theme changes not applying?**
- Call window.Invalidate() after theme changes
- Ensure all components use the same theme instance
- Check that custom colors are valid NRGBA values

**Q: Input components not receiving keyboard input?**
- Ensure the input has focus
- Check that no other component is capturing events
- Verify your event loop is processing key events

## License

MIT License - see LICENSE file for details

## Credits

- Inspired by [shadcn/ui](https://ui.shadcn.com/)
- Built with [Gio](https://gioui.org/)
- Thanks to all contributors

## Support

- [GitHub Issues](https://github.com/bnema/gio-shadcn/issues)
- [Documentation](https://github.com/bnema/gio-shadcn/wiki)
- [Examples](https://github.com/bnema/gio-shadcn/tree/main/examples)