package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Registry represents the component registry structure.
type Registry struct {
	Version    string               `json:"version"`
	Components map[string]Component `json:"components"`
	Categories map[string]Category  `json:"categories"`
}

// Component represents a single UI component.
type Component struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Files        []string `json:"files"`
	Dependencies []string `json:"dependencies"`
	Imports      []string `json:"imports"`
	Version      string   `json:"version"`
	Category     string   `json:"category"`
}

// Category represents a component category.
type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available components",
	Long: `List all available components in the gio-shadcn registry.

This command shows:
- Component names and descriptions
- Categories
- Version information
- Dependencies`,
	RunE: func(_ *cobra.Command, _ []string) error {
		return listComponents()
	},
}

func listComponents() error {
	registry, err := loadRegistry()
	if err != nil {
		return fmt.Errorf("failed to load registry: %w", err)
	}

	fmt.Printf("📦 gio-shadcn components (v%s)\n\n", registry.Version)

	// Group components by category
	componentsByCategory := make(map[string][]Component)
	for _, component := range registry.Components {
		category := component.Category
		if category == "" {
			category = "other"
		}
		componentsByCategory[category] = append(componentsByCategory[category], component)
	}

	// Display components grouped by category
	for categoryKey, components := range componentsByCategory {
		if len(components) == 0 {
			continue
		}

		// Get category info
		var categoryName, categoryDesc string
		if cat, exists := registry.Categories[categoryKey]; exists {
			categoryName = cat.Name
			categoryDesc = cat.Description
		} else {
			categoryName = cases.Title(language.English).String(categoryKey)
			categoryDesc = fmt.Sprintf("%s components", categoryName)
		}

		fmt.Printf("🏷️  %s\n", categoryName)
		fmt.Printf("   %s\n\n", categoryDesc)

		for _, component := range components {
			fmt.Printf("   %-15s %s\n", component.Name, component.Description)
			if len(component.Dependencies) > 0 {
				fmt.Printf("   %-15s Dependencies: %s\n", "", strings.Join(component.Dependencies, ", "))
			}
		}
		fmt.Println()
	}

	fmt.Printf("💡 Usage:\n")
	fmt.Printf("   gio-shadcn add <component-name>\n")
	fmt.Printf("   gio-shadcn add button\n")

	return nil
}

func loadRegistry() (*Registry, error) {
	// First try to load from local file
	if data, err := os.ReadFile("registry.json"); err == nil {
		var registry Registry
		if err := json.Unmarshal(data, &registry); err == nil {
			return &registry, nil
		}
	}

	// If local file doesn't exist, try to fetch from remote
	return fetchRemoteRegistry()
}

func fetchRemoteRegistry() (*Registry, error) {
	url := "https://raw.githubusercontent.com/bnema/gio-shadcn/main/registry.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch remote registry: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Warning: failed to close response body: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch registry: HTTP %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var registry Registry
	if err := json.Unmarshal(data, &registry); err != nil {
		return nil, fmt.Errorf("failed to parse registry: %w", err)
	}

	return &registry, nil
}
