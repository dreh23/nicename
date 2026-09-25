package main

import (
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func TestCLIExecution(t *testing.T) {
	tmpDir := t.TempDir()
	binPath := tmpDir + "/nicename_test_bin"

	buildCmd := exec.Command("go", "build", "-o", binPath, ".")
	if out, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build CLI binary: %v\nOutput: %s", err, string(out))
	}

	taskIDPattern := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)+-[0-9a-f]{4}$`)
	slugPattern := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)+$`)

	t.Run("default output", func(t *testing.T) {
		cmd := exec.Command(binPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Command failed: %v", err)
		}
		res := strings.TrimSpace(string(out))
		if !strings.Contains(res, " ") {
			t.Errorf("Expected space-separated words, got %q", res)
		}
	})

	t.Run("slug flag", func(t *testing.T) {
		cmd := exec.Command(binPath, "-slug")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Command failed: %v", err)
		}
		res := strings.TrimSpace(string(out))
		if !slugPattern.MatchString(res) {
			t.Errorf("Expected slug pattern, got %q", res)
		}
	})

	t.Run("task-id flag", func(t *testing.T) {
		cmd := exec.Command(binPath, "-task-id")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Command failed: %v", err)
		}
		res := strings.TrimSpace(string(out))
		if !taskIDPattern.MatchString(res) {
			t.Errorf("Expected taskID pattern, got %q", res)
		}
	})

	t.Run("classic and task-id combination", func(t *testing.T) {
		cmd := exec.Command(binPath, "-classic", "-task-id")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Command failed: %v", err)
		}
		res := strings.TrimSpace(string(out))
		if !taskIDPattern.MatchString(res) {
			t.Errorf("Expected taskID pattern, got %q", res)
		}
	})

	t.Run("character and task-id combination", func(t *testing.T) {
		cmd := exec.Command(binPath, "-character", "-task-id")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Command failed: %v", err)
		}
		res := strings.TrimSpace(string(out))
		if !taskIDPattern.MatchString(res) {
			t.Errorf("Expected taskID pattern, got %q", res)
		}
	})

	t.Run("count flag", func(t *testing.T) {
		cmd := exec.Command(binPath, "-count", "5")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Command failed: %v", err)
		}
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) != 5 {
			t.Errorf("Expected 5 lines, got %d", len(lines))
		}
	})

	t.Run("count zero", func(t *testing.T) {
		cmd := exec.Command(binPath, "-count", "0")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Command failed: %v", err)
		}
		res := strings.TrimSpace(string(out))
		if res != "" {
			t.Errorf("Expected empty output for count 0, got %q", res)
		}
	})

	t.Run("negative count error", func(t *testing.T) {
		cmd := exec.Command(binPath, "-count", "-1")
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("Expected error for negative count, got success with output %q", string(out))
		}
	})

	t.Run("conflicting classic and character error", func(t *testing.T) {
		cmd := exec.Command(binPath, "-classic", "-character")
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("Expected error for conflicting flags, got success with output %q", string(out))
		}
	})

	t.Run("help flag", func(t *testing.T) {
		cmd := exec.Command(binPath, "-help")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Expected exit code 0 for help flag, got error: %v", err)
		}
		if !strings.Contains(string(out), "Usage: nicename") {
			t.Errorf("Expected help output, got %q", string(out))
		}
	})
}
