package javascript

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// PrettierConfig represents the configuration for Prettier
type PrettierConfig struct {
	Semi          bool   `json:"semi"`
	SingleQuote   bool   `json:"singleQuote"`
	TabWidth      int    `json:"tabWidth"`
	PrintWidth    int    `json:"printWidth"`
	TrailingComma string `json:"trailingComma"`
	BracketSpacing bool   `json:"bracketSpacing"`
}

// DefaultConfig returns the default Prettier configuration
func DefaultConfig() PrettierConfig {
	return PrettierConfig{
		Semi:           true,
		SingleQuote:    false,
		TabWidth:       2,
		PrintWidth:     80,
		TrailingComma:  "es5",
		BracketSpacing: true,
	}
}

// Format formats JavaScript code using Prettier
func Format(content string) (string, error) {
	// Create a temporary directory for our work
	tmpDir, err := os.MkdirTemp("", "kram-js-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Write the content to a temporary file
	inputFile := filepath.Join(tmpDir, "input.js")
	if err := os.WriteFile(inputFile, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write input file: %v", err)
	}

	// Write prettier config
	config := DefaultConfig()
	configBytes, err := json.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config: %v", err)
	}

	configFile := filepath.Join(tmpDir, ".prettierrc")
	if err := os.WriteFile(configFile, configBytes, 0644); err != nil {
		return "", fmt.Errorf("failed to write config file: %v", err)
	}

	// Initialize npm project and install prettier
	cmd := exec.Command("npm", "init", "-y")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to initialize npm project: %v", err)
	}

	cmd = exec.Command("npm", "install", "prettier", "--save-dev")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to install prettier: %v", err)
	}

	// Run prettier
	cmd = exec.Command("npx", "prettier", "--write", "input.js")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run prettier: %v", err)
	}

	// Read the formatted output
	formatted, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read formatted file: %v", err)
	}

	return string(formatted), nil
}
