package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCLIExecution(t *testing.T) {
	// Build temporary binary
	tmpDir := t.TempDir()
	binPath := tmpDir + "/nicename_test_bin"

	buildCmd := exec.Command("go", "build", "-o", binPath, ".")
	if out, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build CLI binary: %v\nOutput: %s", err, string(out))
	}

	tests := []struct {
		name string
		args []string
	}{
		{"default", nil},
		{"slug", []string{"-slug"}},
		{"task-id", []string{"-task-id"}},
		{"classic", []string{"-classic"}},
		{"character", []string{"-character"}},
		{"multi-count", []string{"-count", "3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binPath, tt.args...)
			cmd.Env = os.Environ()
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Command failed: %v\nOutput: %s", err, string(out))
			}
			trimmed := strings.TrimSpace(string(out))
			if trimmed == "" {
				t.Errorf("Expected output, got empty string")
			}
		})
	}
}
