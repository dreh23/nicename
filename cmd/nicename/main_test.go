package main

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	taskIDPattern := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)+-[0-9a-f]{4}$`)
	slugPattern := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)+$`)

	tests := []struct {
		name         string
		args         []string
		expectedCode int
		validate     func(t *testing.T, stdout, stderr string)
	}{
		{
			name:         "default output",
			args:         nil,
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !strings.Contains(res, " ") {
					t.Errorf("Expected space-separated words, got %q", res)
				}
			},
		},
		{
			name:         "slug flag",
			args:         []string{"-slug"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !slugPattern.MatchString(res) {
					t.Errorf("Expected slug pattern, got %q", res)
				}
			},
		},
		{
			name:         "task-id flag",
			args:         []string{"-task-id"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !taskIDPattern.MatchString(res) {
					t.Errorf("Expected taskID pattern, got %q", res)
				}
			},
		},
		{
			name:         "classic flag",
			args:         []string{"-classic"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !strings.Contains(res, " ") {
					t.Errorf("Expected space-separated classic pair, got %q", res)
				}
			},
		},
		{
			name:         "classic and slug combination",
			args:         []string{"-classic", "-slug"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !slugPattern.MatchString(res) {
					t.Errorf("Expected slug pattern, got %q", res)
				}
			},
		},
		{
			name:         "classic and task-id combination",
			args:         []string{"-classic", "-task-id"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !taskIDPattern.MatchString(res) {
					t.Errorf("Expected taskID pattern, got %q", res)
				}
			},
		},
		{
			name:         "character flag",
			args:         []string{"-character"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !strings.Contains(res, " ") {
					t.Errorf("Expected space-separated character pair, got %q", res)
				}
			},
		},
		{
			name:         "character and slug combination",
			args:         []string{"-character", "-slug"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !slugPattern.MatchString(res) {
					t.Errorf("Expected slug pattern, got %q", res)
				}
			},
		},
		{
			name:         "character and task-id combination",
			args:         []string{"-character", "-task-id"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if !taskIDPattern.MatchString(res) {
					t.Errorf("Expected taskID pattern, got %q", res)
				}
			},
		},
		{
			name:         "count flag",
			args:         []string{"-count", "3"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				lines := strings.Split(strings.TrimSpace(stdout), "\n")
				if len(lines) != 3 {
					t.Errorf("Expected 3 lines, got %d", len(lines))
				}
			},
		},
		{
			name:         "count zero",
			args:         []string{"-count", "0"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				res := strings.TrimSpace(stdout)
				if res != "" {
					t.Errorf("Expected empty output for count 0, got %q", res)
				}
			},
		},
		{
			name:         "negative count error",
			args:         []string{"-count", "-5"},
			expectedCode: 1,
			validate: func(t *testing.T, stdout, stderr string) {
				if !strings.Contains(stderr, "Error: -count must be non-negative") {
					t.Errorf("Expected negative count error, got %q", stderr)
				}
			},
		},
		{
			name:         "conflicting classic and character error",
			args:         []string{"-classic", "-character"},
			expectedCode: 1,
			validate: func(t *testing.T, stdout, stderr string) {
				if !strings.Contains(stderr, "Error: cannot combine -classic and -character") {
					t.Errorf("Expected conflicting flags error, got %q", stderr)
				}
			},
		},
		{
			name:         "help flag",
			args:         []string{"-help"},
			expectedCode: 0,
			validate: func(t *testing.T, stdout, stderr string) {
				if !strings.Contains(stderr, "Usage: nicename") {
					t.Errorf("Expected help output in stderr, got %q", stderr)
				}
			},
		},
		{
			name:         "unknown flag error",
			args:         []string{"-unknown-flag"},
			expectedCode: 2,
			validate: func(t *testing.T, stdout, stderr string) {
				if !strings.Contains(stderr, "flag provided but not defined") {
					t.Errorf("Expected unknown flag error, got %q", stderr)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, &stdout, &stderr)
			if code != tt.expectedCode {
				t.Errorf("run(%v) exit code = %d; want %d", tt.args, code, tt.expectedCode)
			}
			if tt.validate != nil {
				tt.validate(t, stdout.String(), stderr.String())
			}
		})
	}
}

func TestBinarySubprocess(t *testing.T) {
	tmpDir := t.TempDir()
	binPath := tmpDir + "/nicename_bin"

	buildCmd := exec.Command("go", "build", "-o", binPath, ".")
	if out, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build CLI binary: %v\nOutput: %s", err, string(out))
	}

	cmd := exec.Command(binPath, "-slug")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Binary execution failed: %v\nOutput: %s", err, string(out))
	}
	if strings.TrimSpace(string(out)) == "" {
		t.Errorf("Expected non-empty output from binary")
	}
}

func TestMainFunction(t *testing.T) {
	origExit := osExit
	origArgs := os.Args
	defer func() {
		osExit = origExit
		os.Args = origArgs
	}()

	var capturedExit int
	osExit = func(code int) {
		capturedExit = code
	}

	os.Args = []string{"nicename", "-slug"}
	main()

	if capturedExit != 0 {
		t.Errorf("main() exited with code %d; want 0", capturedExit)
	}
}
