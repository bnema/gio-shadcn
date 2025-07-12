/*
Package gio-shadcn provides a comprehensive collection of UI components for Gio applications,
inspired by shadcn/ui design principles and patterns.

# Overview

gio-shadcn is a Go port of the popular shadcn/ui component library, bringing beautiful,
accessible, and customizable UI components to Gio applications. Each component follows
consistent design patterns and integrates seamlessly with a powerful theming system.

# Installation

Add gio-shadcn to your Go project:

	go get github.com/bnema/gio-shadcn@latest

# Quick Start

Here's a minimal example showing how to use gio-shadcn components:

	package main

	import (
		"gioui.org/app"
		"gioui.org/io/system"
		"gioui.org/layout"
		"gioui.org/op"

		"github.com/bnema/gio-shadcn/components/button"
		"github.com/bnema/gio-shadcn/theme"
	)

	func main() {
		go func() {
			w := app.NewWindow(app.Title("gio-shadcn Example"))
			defer w.Close()

			// Create theme
			th := theme.New()

			// Create button with functional options
			btn := button.New(button.Config{
				Text:    "Click me",
				Variant: theme.VariantDefault,
				OnClick: func() {
					println("Button clicked!")
				},
			})

			var ops op.Ops
			for {
				switch e := w.Event().(type) {
				case system.DestroyEvent:
					return
				case system.FrameEvent:
					gtx := app.NewContext(&ops, e)

					// Layout the button
					layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return btn.Layout(gtx, th)
					})

					e.Frame(&ops)
				}
			}
		}()
		app.Main()
	}

# Components

gio-shadcn provides the following components (5 of 51 planned):

• Button - Versatile button with multiple variants (default, destructive, outline, secondary, ghost, link)
• Card - Container component for grouping related content
• Input - Text input with validation and different types
• Label - Typography component with semantic variants
• Titlebar - Custom window titlebar for frameless windows

# Theme System

The theme system provides comprehensive control over your application's appearance:

	// Create and customize theme
	th := theme.New()

	// Toggle between light and dark mode
	th.ToggleDark()

	// Load theme from JSON file
	th, err := theme.NewFromJSON("theme.json")

# Architecture

All components follow consistent patterns:

1. Creation: Use functional options pattern for configuration
2. Layout: Call Layout(gtx, theme) to render the component
3. State: Components manage their own internal state
4. Theming: All components respect the current theme

# Import Paths

Components:

	github.com/bnema/gio-shadcn/components/button
	github.com/bnema/gio-shadcn/components/card
	github.com/bnema/gio-shadcn/components/input
	github.com/bnema/gio-shadcn/components/label
	github.com/bnema/gio-shadcn/components/titlebar

Theme system:

	github.com/bnema/gio-shadcn/theme

Utilities:

	github.com/bnema/gio-shadcn/utils

# Examples

See the demo application for comprehensive examples:

	go run ./cmd/demo-app

# Development Status

This project is under active development. While functional, some features are still
being implemented. Currently 5 of 51 planned components are available.

# License

MIT License - see LICENSE file for details.
*/
package main
