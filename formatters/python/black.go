package python

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// BlackConfig represents the configuration for Black formatter
type BlackConfig struct {
	LineLength  int  `json:"line_length"`
	SkipString bool `json:"skip_string_normalization"`
	FastMode   bool `json:"fast"`
}

// DefaultConfig returns the default Black configuration
func DefaultConfig() BlackConfig {
	return BlackConfig{
		LineLength:  88,  // Black's default
		SkipString: false,
		FastMode:   false,
	}
}

// Format formats Python code using Black
func Format(content string) (string, error) {
	// Check if Python is installed
	pythonPath, err := exec.LookPath("python3")
	if err != nil {
		return "", fmt.Errorf("python3 not found: %v", err)
	}

	// Create a temporary directory for our work
	tmpDir, err := os.MkdirTemp("", "kram-py-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Write the content to a temporary file
	inputFile := filepath.Join(tmpDir, "input.py")
	if err := os.WriteFile(inputFile, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write input file: %v", err)
	}

	// Create a virtual environment
	cmd := exec.Command(pythonPath, "-m", "venv", "venv")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to create virtual environment: %v", err)
	}

	// Install black in the virtual environment
	pipCmd := filepath.Join(tmpDir, "venv", "bin", "pip")
	cmd = exec.Command(pipCmd, "install", "black")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to install black: %v", err)
	}

	// Run black
	blackCmd := filepath.Join(tmpDir, "venv", "bin", "black")
	cmd = exec.Command(blackCmd, "-q", inputFile)
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run black: %v\nThis might indicate a syntax error in the Python code", err)
	}

	// Read the formatted output
	formatted, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read formatted file: %v", err)
	}

	return string(formatted), nil
}
