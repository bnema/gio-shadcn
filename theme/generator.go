package theme

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"os"
	"strconv"
	"strings"
)

// Config represents a complete theme configuration that can be loaded from JSON.
// This struct provides a way to define themes externally using JSON files,.
// allowing for easy theme customization and distribution. It supports both
// light and dark color schemes, custom radius values, and metadata.
//
// Example JSON structure:.
//
//	{
//		"name": "Custom Theme",
//		"description": "A beautiful custom theme",
//		"version": "1.0.0",
//		"colors": {
//			"light": {
//				"background": "#ffffff",
//				"foreground": "#000000",
//				"primary": "#3b82f6"
//			},
//			"dark": {
//				"background": "#000000",
//				"foreground": "#ffffff",
//				"primary": "#60a5fa"
//			}
//		},
//		"radius": {
//			"sm": "2px",
//			"md": "4px",
//			"lg": "8px"
//		}
//	}
type Config struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Colors      struct {
		Light map[string]string `json:"light"`
		Dark  map[string]string `json:"dark"`
	} `json:"colors"`
	Radius struct {
		None string `json:"none"`
		SM   string `json:"sm"`
		MD   string `json:"md"`
		LG   string `json:"lg"`
		XL   string `json:"xl"`
		XXL  string `json:"2xl"`
		XXXL string `json:"3xl"`
		Full string `json:"full"`
	} `json:"radius"`
	Spacing    map[string]int `json:"spacing"`
	Typography struct {
		FontFamily map[string][]string `json:"fontFamily"`
		FontSize   map[string]int      `json:"fontSize"`
		FontWeight map[string]int      `json:"fontWeight"`
		LineHeight map[string]float64  `json:"lineHeight"`
	} `json:"typography"`
}

// LoadThemeFromJSON loads a theme configuration from a JSON file.
// This function reads and parses a JSON theme file, returning a Config.
// struct that can be used to generate Theme instances. The JSON should follow
// the shadcn/ui theme format with light/dark color schemes.
//
// Example usage:.
//
//	config, err := LoadThemeFromJSON("mytheme.json")
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Use config to create theme
func LoadThemeFromJSON(path string) (*Config, error) {
	//nolint:gosec // File path is intended to be user-provided for theme loading
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open theme file: %w", err)
	}
	defer func() { _ = file.Close() }()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read theme file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse theme JSON: %w", err)
	}

	return &config, nil
}

func hexToNRGBA(hex string) (color.NRGBA, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return color.NRGBA{}, fmt.Errorf("invalid hex color format: %s", hex)
	}

	r, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return color.NRGBA{}, err
	}

	g, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return color.NRGBA{}, err
	}

	b, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return color.NRGBA{}, err
	}

	return color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}, nil
}

// ToColorScheme converts a Config to a ColorScheme for the specified mode.
// This method extracts color definitions from the JSON configuration and.
// converts them to the internal ColorScheme format used by gio-shadcn themes.
// Set isDark to true for dark mode colors, false for light mode colors.
//
// Returns an error if any required colors are missing or invalid hex values.
func (config *Config) ToColorScheme(isDark bool) (ColorScheme, error) {
	var colorMap map[string]string
	if isDark {
		colorMap = config.Colors.Dark
	} else {
		colorMap = config.Colors.Light
	}

	cs := ColorScheme{}

	// Helper function to convert hex to NRGBA and handle errors
	convertColor := func(name, hex string) (color.NRGBA, error) {
		if hex == "" {
			return color.NRGBA{}, fmt.Errorf("missing color for %s", name)
		}
		return hexToNRGBA(hex)
	}

	var err error
	cs.Background, err = convertColor("background", colorMap["background"])
	if err != nil {
		return cs, err
	}

	cs.Foreground, err = convertColor("foreground", colorMap["foreground"])
	if err != nil {
		return cs, err
	}

	cs.Card, err = convertColor("card", colorMap["card"])
	if err != nil {
		return cs, err
	}

	cs.CardFg, err = convertColor("card-foreground", colorMap["card-foreground"])
	if err != nil {
		return cs, err
	}

	cs.Popover, err = convertColor("popover", colorMap["popover"])
	if err != nil {
		return cs, err
	}

	cs.PopoverFg, err = convertColor("popover-foreground", colorMap["popover-foreground"])
	if err != nil {
		return cs, err
	}

	cs.Primary, err = convertColor("primary", colorMap["primary"])
	if err != nil {
		return cs, err
	}

	cs.PrimaryFg, err = convertColor("primary-foreground", colorMap["primary-foreground"])
	if err != nil {
		return cs, err
	}

	cs.Secondary, err = convertColor("secondary", colorMap["secondary"])
	if err != nil {
		return cs, err
	}

	cs.SecondaryFg, err = convertColor("secondary-foreground", colorMap["secondary-foreground"])
	if err != nil {
		return cs, err
	}

	cs.Muted, err = convertColor("muted", colorMap["muted"])
	if err != nil {
		return cs, err
	}

	cs.MutedFg, err = convertColor("muted-foreground", colorMap["muted-foreground"])
	if err != nil {
		return cs, err
	}

	cs.Accent, err = convertColor("accent", colorMap["accent"])
	if err != nil {
		return cs, err
	}

	cs.AccentFg, err = convertColor("accent-foreground", colorMap["accent-foreground"])
	if err != nil {
		return cs, err
	}

	cs.Destructive, err = convertColor("destructive", colorMap["destructive"])
	if err != nil {
		return cs, err
	}

	cs.DestructiveFg, err = convertColor("destructive-foreground", colorMap["destructive-foreground"])
	if err != nil {
		return cs, err
	}

	cs.Border, err = convertColor("border", colorMap["border"])
	if err != nil {
		return cs, err
	}

	cs.Input, err = convertColor("input", colorMap["input"])
	if err != nil {
		return cs, err
	}

	cs.Ring, err = convertColor("ring", colorMap["ring"])
	if err != nil {
		return cs, err
	}

	return cs, nil
}

// NewThemeFromJSON creates a complete Theme instance from a JSON configuration file.
// This is the main function for loading external themes. It loads the JSON config,
// converts both light and dark color schemes, and creates a fully functional Theme.
// with default typography and spacing. The resulting theme can be used immediately
// with all gio-shadcn components.
//
// Example usage:.
//
//	theme, err := NewThemeFromJSON("themes/custom.json")
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Use theme with components
//	button.Layout(gtx, theme)
func NewThemeFromJSON(path string) (*Theme, error) {
	config, err := LoadThemeFromJSON(path)
	if err != nil {
		return nil, err
	}

	lightColors, err := config.ToColorScheme(false)
	if err != nil {
		return nil, fmt.Errorf("failed to parse light colors: %w", err)
	}

	darkColors, err := config.ToColorScheme(true)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dark colors: %w", err)
	}

	return &Theme{
		Colors:     lightColors,
		DarkColors: darkColors,
		Typography: DefaultTypography(),
		Spacing:    DefaultSpacing(),
		IsDark:     false,
	}, nil
}

// GenerateThemeConstants generates Go source code with color constants from a theme config.
// This function creates Go constants for all colors in both light and dark themes,.
// which can be used for code generation or creating static theme definitions.
// The generated code includes proper imports and follows Go naming conventions.
//
// Returns a string containing complete Go source code with color constants.
func GenerateThemeConstants(config *Config) string {
	var sb strings.Builder

	sb.WriteString("// Code generated by theme generator. DO NOT EDIT.\n\n")
	sb.WriteString("package theme\n\n")
	sb.WriteString("import \"image/color\"\n\n")

	// Generate light theme constants
	sb.WriteString("// Light theme colors\n")
	sb.WriteString("var (\n")
	for name, hex := range config.Colors.Light {
		constantName := strings.ToUpper(strings.ReplaceAll(name, "-", "_")) + "_LIGHT"
		rgba, _ := hexToNRGBA(hex)
		sb.WriteString(fmt.Sprintf("\t%s = color.NRGBA{R: %d, G: %d, B: %d, A: %d}\n",
			constantName, rgba.R, rgba.G, rgba.B, rgba.A))
	}
	sb.WriteString(")\n\n")

	// Generate dark theme constants
	sb.WriteString("// Dark theme colors\n")
	sb.WriteString("var (\n")
	for name, hex := range config.Colors.Dark {
		constantName := strings.ToUpper(strings.ReplaceAll(name, "-", "_")) + "_DARK"
		rgba, _ := hexToNRGBA(hex)
		sb.WriteString(fmt.Sprintf("\t%s = color.NRGBA{R: %d, G: %d, B: %d, A: %d}\n",
			constantName, rgba.R, rgba.G, rgba.B, rgba.A))
	}
	sb.WriteString(")\n\n")

	return sb.String()
}
