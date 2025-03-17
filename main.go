package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/divyangchauhan/kram/formatters/javascript"
	"github.com/divyangchauhan/kram/formatters/python"
)

// Formatter defines the interface for language-specific formatters
type Formatter interface {
	Format(content string) (string, error)
}

// PythonFormatter implements Formatter for Python
type PythonFormatter struct{}

func (f *PythonFormatter) Format(content string) (string, error) {
	return python.Format(content)
}

// JavaScriptFormatter implements Formatter for JavaScript
type JavaScriptFormatter struct{}

func (f *JavaScriptFormatter) Format(content string) (string, error) {
	// Check if Node.js is installed
	if _, err := exec.LookPath("node"); err != nil {
		return "", fmt.Errorf("Node.js is required for JavaScript formatting but not found: %v", err)
	}

	// Check if npm is installed
	if _, err := exec.LookPath("npm"); err != nil {
		return "", fmt.Errorf("npm is required for JavaScript formatting but not found: %v", err)
	}

	return javascript.Format(content)
}

// GoFormatter implements Formatter for Go
type GoFormatter struct{}

func (f *GoFormatter) Format(content string) (string, error) {
	return "", fmt.Errorf("Go formatting not yet implemented")
}

// JavaFormatter implements Formatter for Java
type JavaFormatter struct{}

func (f *JavaFormatter) Format(content string) (string, error) {
	return "", fmt.Errorf("Java formatting not yet implemented")
}

// getFormatter returns the appropriate formatter based on file extension
func getFormatter(filename string) Formatter {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".py":
		return &PythonFormatter{}
	case ".js":
		return &JavaScriptFormatter{}
	case ".go":
		return &GoFormatter{}
	case ".java":
		return &JavaFormatter{}
	default:
		return nil
	}
}

func main() {
	recursive := flag.Bool("r", false, "Format files recursively")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Error: No files specified")
		os.Exit(1)
	}

	for _, path := range files {
		if *recursive {
			err := filepath.Walk(path, processFile)
			if err != nil {
				fmt.Printf("Error walking path %s: %v\n", path, err)
			}
		} else {
			if err := processFile(path, nil, nil); err != nil {
				fmt.Printf("Error processing file %s: %v\n", path, err)
			}
		}
	}
}

func processFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// Skip directories when processing individual files
	if info != nil && info.IsDir() {
		return nil
	}

	// Get the formatter based on file extension
	formatter := getFormatter(path)
	if formatter == nil {
		return fmt.Errorf("unsupported file type: %s", path)
	}

	// Read the file
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Format the content
	formatted, err := formatter.Format(string(content))
	if err != nil {
		return fmt.Errorf("error formatting file: %v", err)
	}

	// Write back to file
	err = os.WriteFile(path, []byte(formatted), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	fmt.Printf("Formatted %s\n", path)
	return nil
}
