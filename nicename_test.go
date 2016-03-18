package nicename

import (
	"testing"
)

func TestReturnPairEmpty(t *testing.T) {
	result := GeneratePair()
	if result == "" {
		t.Error("Expected String, got ", result)
	}
}
