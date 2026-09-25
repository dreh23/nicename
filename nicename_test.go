package nicename

import (
	"errors"
	"regexp"
	"strings"
	"sync"
	"testing"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)+$`)

func TestGeneratePair(t *testing.T) {
	pair := GeneratePair()
	if pair == "" {
		t.Fatal("GeneratePair returned empty string")
	}
	parts := strings.Split(pair, " ")
	if len(parts) < 2 {
		t.Fatalf("Expected at least two words in pair %q", pair)
	}

	// Test empty slice fallback
	origFirst := First
	First = nil
	if empty := GeneratePair(); empty != "" {
		t.Errorf("Expected empty string when First is nil, got %q", empty)
	}
	First = origFirst
}

func TestGeneratePairSlug(t *testing.T) {
	slug := GeneratePairSlug()
	if slug == "" {
		t.Fatal("GeneratePairSlug returned empty string")
	}
	if !slugRegex.MatchString(slug) {
		t.Errorf("GeneratePairSlug %q does not match slug pattern", slug)
	}
}

func TestGenerateAardman(t *testing.T) {
	val := GenerateAardman()
	if val == "" {
		t.Fatal("GenerateAardman returned empty string")
	}
	parts := strings.Split(val, " ")
	if len(parts) < 2 {
		t.Fatalf("Expected at least two words, got %q", val)
	}

	// Test empty list handling
	origAdj := AardmanAdjectives
	AardmanAdjectives = nil
	if empty := GenerateAardman(); empty != "" {
		t.Errorf("Expected empty string when AardmanAdjectives is nil, got %q", empty)
	}
	AardmanAdjectives = origAdj
}

func TestGenerateAardmanCharacter(t *testing.T) {
	val := GenerateAardmanCharacter()
	if val == "" {
		t.Fatal("GenerateAardmanCharacter returned empty string")
	}

	origChars := AardmanCharacters
	AardmanCharacters = nil
	if empty := GenerateAardmanCharacter(); empty != "" {
		t.Errorf("Expected empty string when AardmanCharacters is nil, got %q", empty)
	}
	AardmanCharacters = origChars
}

func TestGenerateAardmanSlug(t *testing.T) {
	for i := 0; i < 50; i++ {
		slug := GenerateAardmanSlug()
		if slug == "" {
			t.Fatal("GenerateAardmanSlug returned empty string")
		}
		if !slugRegex.MatchString(slug) {
			t.Errorf("GenerateAardmanSlug %q does not match slug format", slug)
		}
		if strings.Contains(slug, " ") || strings.Contains(slug, "_") {
			t.Errorf("Slug %q contains invalid characters", slug)
		}
	}
}

func TestGenerateTaskID(t *testing.T) {
	taskIDRegex := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)+-[0-9a-f]{4}$`)
	seen := make(map[string]bool)

	for i := 0; i < 100; i++ {
		taskID := GenerateTaskID()
		if !taskIDRegex.MatchString(taskID) {
			t.Errorf("GenerateTaskID %q does not match pattern [slug]-[hex4]", taskID)
		}
		if seen[taskID] {
			t.Errorf("Collision detected in 100 iterations: %q", taskID)
		}
		seen[taskID] = true
	}
}

func TestFormatTaskID(t *testing.T) {
	hex4Regex := regexp.MustCompile(`^[0-9a-f]{4}$`)
	// Empty slug should return only 4 hex characters without leading hyphen
	emptyResult := FormatTaskID("")
	if !hex4Regex.MatchString(emptyResult) {
		t.Errorf("FormatTaskID(\"\") = %q; want 4 hex chars", emptyResult)
	}

	customResult := FormatTaskID("build-task")
	if !strings.HasPrefix(customResult, "build-task-") || len(customResult) != len("build-task-")+4 {
		t.Errorf("FormatTaskID(\"build-task\") = %q; want prefix build-task-[hex4]", customResult)
	}
}

func TestFormatTaskID_EntropyFallback(t *testing.T) {
	origReader := cryptoRandReader
	defer func() { cryptoRandReader = origReader }()

	cryptoRandReader = func(b []byte) (int, error) {
		return 0, errors.New("simulated entropy failure")
	}

	res := FormatTaskID("worker-task")
	if !strings.HasPrefix(res, "worker-task-") || len(res) != len("worker-task-")+4 {
		t.Errorf("FormatTaskID fallback failed, got %q", res)
	}

	emptyRes := FormatTaskID("")
	if len(emptyRes) != 4 || strings.HasPrefix(emptyRes, "-") {
		t.Errorf("FormatTaskID fallback with empty slug failed, got %q", emptyRes)
	}
}

func TestGenerateCustom(t *testing.T) {
	adjs := []string{"Super", "Mega"}
	nouns := []string{"Widget", "Contraption"}

	custom := GenerateCustom(adjs, nouns, " ", false)
	if custom == "" {
		t.Fatal("Expected custom result, got empty")
	}

	slug := GenerateCustom(adjs, nouns, " ", true)
	if !slugRegex.MatchString(slug) {
		t.Errorf("Expected slug format, got %q", slug)
	}

	// Empty slices
	if empty := GenerateCustom(nil, nouns, "-", true); empty != "" {
		t.Errorf("Expected empty for nil adjs, got %q", empty)
	}
	if empty := GenerateCustom(adjs, nil, "-", true); empty != "" {
		t.Errorf("Expected empty for nil nouns, got %q", empty)
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Feathers McGraw", "feathers-mcgraw"},
		{"  Cracking   Toast!  ", "cracking-toast"},
		{"Bun-vac 6000", "bun-vac-6000"},
		{"62 West Wallaby Street", "62-west-wallaby-street"},
		{"Anti-Pesto Van!", "anti-pesto-van"},
		{"---hello---world---", "hello-world"},
		{"Wallace & Gromit", "wallace-gromit"},
		{"", ""},
		{"   ", ""},
		{"---", ""},
		{"!@#$%^&*()", ""},
		{" - - hello - - ", "hello"},
		{"12345", "12345"},
		{"foo_bar_baz", "foo-bar-baz"},
	}

	for _, tt := range tests {
		got := Slugify(tt.input)
		if got != tt.expected {
			t.Errorf("Slugify(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				_ = GeneratePair()
				_ = GeneratePairSlug()
				_ = GenerateAardman()
				_ = GenerateAardmanSlug()
				_ = GenerateTaskID()
			}
		}()
	}
	wg.Wait()
}

func BenchmarkGeneratePair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GeneratePair()
	}
}

func BenchmarkGenerateAardmanSlug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateAardmanSlug()
	}
}

func BenchmarkGenerateTaskID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateTaskID()
	}
}

func BenchmarkFormatTaskID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormatTaskID("cracking-gromit")
	}
}
