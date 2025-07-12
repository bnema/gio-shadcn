package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new gio-shadcn project",
	Long: `Initialize a new gio-shadcn project with the basic structure and configuration files.

This command creates:
- go.mod file for the project
- theme.json configuration file
- components directory structure
- basic example files`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		projectName := "my-gio-app"
		if len(args) > 0 {
			projectName = args[0]
		}

		return initProject(projectName)
	},
}

func initProject(projectName string) error {
	// Create project directory
	if err := os.MkdirAll(projectName, 0750); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create go.mod
	goModContent := fmt.Sprintf(`module %s

go 1.24.5

require (
	gioui.org v0.8.0
	github.com/bnema/gio-shadcn/theme v0.0.0
	github.com/bnema/gio-shadcn/utils v0.0.0
)
`, projectName)

	if err := os.WriteFile(filepath.Join(projectName, "go.mod"), []byte(goModContent), 0600); err != nil {
		return fmt.Errorf("failed to create go.mod: %w", err)
	}

	// Create theme.json (copy from the registry)
	themeData := map[string]interface{}{
		"name":        "default",
		"description": "Default theme configuration",
		"version":     "1.0.0",
		"colors": map[string]interface{}{
			"light": map[string]string{
				"background":             "#ffffff",
				"foreground":             "#0a0a0a",
				"primary":                "#0f172a",
				"primary-foreground":     "#f8fafc",
				"secondary":              "#f1f5f9",
				"secondary-foreground":   "#0f172a",
				"muted":                  "#f1f5f9",
				"muted-foreground":       "#64748b",
				"accent":                 "#f1f5f9",
				"accent-foreground":      "#0f172a",
				"destructive":            "#ef4444",
				"destructive-foreground": "#f8fafc",
				"border":                 "#e2e8f0",
				"input":                  "#e2e8f0",
				"ring":                   "#0f172a",
			},
			"dark": map[string]string{
				"background":             "#0a0a0a",
				"foreground":             "#fafafa",
				"primary":                "#fafafa",
				"primary-foreground":     "#0a0a0a",
				"secondary":              "#262626",
				"secondary-foreground":   "#fafafa",
				"muted":                  "#262626",
				"muted-foreground":       "#a1a1aa",
				"accent":                 "#262626",
				"accent-foreground":      "#fafafa",
				"destructive":            "#7f1d1d",
				"destructive-foreground": "#fafafa",
				"border":                 "#262626",
				"input":                  "#262626",
				"ring":                   "#d4d4d8",
			},
		},
		"radius": map[string]interface{}{
			"none": 0,
			"sm":   2,
			"md":   4,
			"lg":   8,
			"xl":   12,
			"2xl":  16,
			"3xl":  24,
			"full": 9999,
		},
	}

	themeJSON, err := json.MarshalIndent(themeData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal theme data: %w", err)
	}

	if err := os.WriteFile(filepath.Join(projectName, "theme.json"), themeJSON, 0600); err != nil {
		return fmt.Errorf("failed to create theme.json: %w", err)
	}

	// Create components directory
	if err := os.MkdirAll(filepath.Join(projectName, "components"), 0750); err != nil {
		return fmt.Errorf("failed to create components directory: %w", err)
	}

	// Create main.go with basic example
	mainGoContent := fmt.Sprintf(`package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"gioui.org/font/gofont"
)

func main() {
	go func() {
		w := app.NewWindow(app.Title("%s"))
		defer w.Close()
		
		th := material.NewTheme()
		th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
		
		var ops op.Ops
		
		for {
			e := w.NextEvent()
			switch e := e.(type) {
			case system.DestroyEvent:
				return
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				
				// Your UI layout goes here
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.H1(th, "Hello, gio-shadcn!").Layout(gtx)
				})
				
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
`, projectName)

	if err := os.WriteFile(filepath.Join(projectName, "main.go"), []byte(mainGoContent), 0600); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}

	fmt.Printf("✅ Successfully initialized gio-shadcn project: %s\n", projectName)
	fmt.Printf("📁 Project structure created:\n")
	fmt.Printf("   %s/\n", projectName)
	fmt.Printf("   ├── go.mod\n")
	fmt.Printf("   ├── theme.json\n")
	fmt.Printf("   ├── components/\n")
	fmt.Printf("   └── main.go\n")
	fmt.Printf("\n🚀 Next steps:\n")
	fmt.Printf("   cd %s\n", projectName)
	fmt.Printf("   gio-shadcn add button\n")
	fmt.Printf("   go run main.go\n")

	return nil
}
