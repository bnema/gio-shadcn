// Package main provides the gio-shadcn CLI tool for managing UI components.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <component-name>",
	Short: "Add a component to your project",
	Long: `Add a component to your project by downloading it from the registry.

This command:
- Downloads the component files
- Installs dependencies if needed
- Rewrites import paths to match your project
- Creates the component directory structure

Examples:
  gio-shadcn add button
  gio-shadcn add card
  gio-shadcn add input`,
	Args: cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		componentName := args[0]
		return addComponent(componentName)
	},
}

func addComponent(componentName string) error {
	// Load registry
	registry, err := loadRegistry()
	if err != nil {
		return fmt.Errorf("failed to load registry: %w", err)
	}

	// Find component
	component, exists := registry.Components[componentName]
	if !exists {
		return fmt.Errorf("component '%s' not found in registry", componentName)
	}

	fmt.Printf("📦 Adding component: %s\n", component.Name)
	fmt.Printf("   Description: %s\n", component.Description)

	// Check if components directory exists
	if _, err := os.Stat("components"); os.IsNotExist(err) {
		if err := os.MkdirAll("components", 0750); err != nil {
			return fmt.Errorf("failed to create components directory: %w", err)
		}
	}

	// Create component directory
	componentDir := filepath.Join("components", componentName)
	if err := os.MkdirAll(componentDir, 0750); err != nil {
		return fmt.Errorf("failed to create component directory: %w", err)
	}

	// Install dependencies first
	if err := installDependencies(component.Dependencies); err != nil {
		return fmt.Errorf("failed to install dependencies: %w", err)
	}

	// Download and install component files
	for _, file := range component.Files {
		if err := downloadComponentFile(componentName, file, componentDir); err != nil {
			return fmt.Errorf("failed to download file %s: %w", file, err)
		}
	}

	// Rewrite import paths
	if err := rewriteImportPaths(componentDir, component.Files); err != nil {
		fmt.Printf("⚠️  Warning: failed to rewrite import paths: %v\n", err)
	}

	fmt.Printf("✅ Successfully added component: %s\n", componentName)
	fmt.Printf("📁 Files installed in: %s\n", componentDir)

	if len(component.Dependencies) > 0 {
		fmt.Printf("📦 Dependencies installed: %s\n", strings.Join(component.Dependencies, ", "))
	}

	return nil
}

func installDependencies(dependencies []string) error {
	for _, dep := range dependencies {
		// Skip if dependency is already installed locally
		if _, err := os.Stat(dep); err == nil {
			continue
		}

		// For theme and utils, we need to download them
		if dep == "theme" || dep == "utils" || strings.HasPrefix(dep, "utils/") {
			if err := downloadDependency(dep); err != nil {
				return fmt.Errorf("failed to download dependency %s: %w", dep, err)
			}
		}
	}
	return nil
}

func downloadDependency(dependency string) error {
	baseURL := "https://raw.githubusercontent.com/bnema/gio-shadcn/main/"

	// Create local directory structure
	if err := os.MkdirAll(dependency, 0750); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dependency, err)
	}

	// Download files based on dependency type
	var filesToDownload []string

	switch dependency {
	case "theme":
		filesToDownload = []string{"theme.go", "colors.go", "typography.go", "spacing.go", "generator.go"}
	case "utils":
		filesToDownload = []string{"cn.go", "variants.go"}
	case "utils/cn":
		filesToDownload = []string{"cn.go"}
	}

	for _, file := range filesToDownload {
		url := fmt.Sprintf("%s%s/%s", baseURL, dependency, file)
		destPath := filepath.Join(dependency, file)

		if err := downloadFile(url, destPath); err != nil {
			return fmt.Errorf("failed to download %s: %w", file, err)
		}
	}

	return nil
}

func downloadComponentFile(componentName, fileName, destDir string) error {
	url := fmt.Sprintf("https://raw.githubusercontent.com/bnema/gio-shadcn/main/components/%s/%s", componentName, fileName)
	destPath := filepath.Join(destDir, fileName)

	return downloadFile(url, destPath)
}

func downloadFile(url, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download from %s: %w", url, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Warning: failed to close response body: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: HTTP %d", resp.StatusCode)
	}

	// Create destination directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(destPath), 0750); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create destination file
	file, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Warning: failed to close file: %v\n", err)
		}
	}()

	// Copy content
	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func rewriteImportPaths(componentDir string, files []string) error {
	// Get current module name from go.mod
	moduleName, err := getCurrentModuleName()
	if err != nil {
		return fmt.Errorf("failed to get module name: %w", err)
	}

	// Rewrite import paths in each file
	for _, file := range files {
		filePath := filepath.Join(componentDir, file)
		if err := rewriteImportsInFile(filePath, moduleName); err != nil {
			return fmt.Errorf("failed to rewrite imports in %s: %w", file, err)
		}
	}

	return nil
}

func getCurrentModuleName() (string, error) {
	goModContent, err := os.ReadFile("go.mod")
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %w", err)
	}

	lines := strings.Split(string(goModContent), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

func rewriteImportsInFile(filePath, moduleName string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Replace import paths
	contentStr := string(content)

	// Replace theme import
	contentStr = strings.ReplaceAll(contentStr,
		`"github.com/bnema/gio-shadcn/theme"`,
		fmt.Sprintf(`"%s/theme"`, moduleName))

	// Replace utils imports
	contentStr = strings.ReplaceAll(contentStr,
		`"github.com/bnema/gio-shadcn/utils"`,
		fmt.Sprintf(`"%s/utils"`, moduleName))

	contentStr = strings.ReplaceAll(contentStr,
		`"github.com/bnema/gio-shadcn/utils/cn"`,
		fmt.Sprintf(`"%s/utils"`, moduleName))

	// Write back to file
	if err := os.WriteFile(filePath, []byte(contentStr), 0600); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
