package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Create a temporary test file for Go (unimplemented formatter)
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.go")
	testContent := "package main\n\nfunc main() {\n\tfmt.Println(\"Hello\")\n}\n"
	
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test processing a single file
	err = processFile(testFile, nil, nil)
	if err == nil {
		t.Error("Expected error for unimplemented Go formatter, got nil")
	}
}

func TestGetFormatter(t *testing.T) {
	tests := []struct {
		filename string
		wantNil  bool
	}{
		{"test.py", false},
		{"test.js", false},
		{"test.go", false},
		{"test.java", false},
		{"test.unknown", true},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got := getFormatter(tt.filename)
			if (got == nil) != tt.wantNil {
				t.Errorf("getFormatter(%q) = %v, want nil: %v", tt.filename, got, tt.wantNil)
			}
		})
	}
}
