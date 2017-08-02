package nows

import (
	"testing"
)

// TestNoWs replaces every w into two v's
func TestNoWs(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name   string
		input  string
		output string
	}{
		{"one w", "paw", "pavv"},
		{"two ws", "wow", "vvovv"},
		{"no w", "hi", "hi"},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			if result := NoWs(tc.input); result != tc.output {
				t.Fatalf("expected %q but got %q", tc.output, result)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name   string
		input  string
		output string
	}{
		{"simple", "hi", "ih"},
		{"empty", "", ""},
		{"arabic", "مروان", "ناورم"},
		{"japanese", "もしもし", "しもしも"},
		{"emojis", "🔥🙏", "🙏🔥"},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			if result := ReverseString(tc.input); result != tc.output {
				t.Fatalf("expected %q but got %q", tc.output, result)
			}
		})
	}
}
