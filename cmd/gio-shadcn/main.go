package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
	rootCmd = &cobra.Command{
		Use:   "gio-shadcn",
		Short: "A CLI tool for managing gio-shadcn components",
		Long: `gio-shadcn is a CLI tool that helps you manage and install
reusable UI components for Gio applications, inspired by shadcn/ui.

This tool allows you to:
- Initialize new projects with base setup
- Add components to your project
- List available components
- Generate themes from configuration files`,
		Version: version,
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(themeCmd)
}
