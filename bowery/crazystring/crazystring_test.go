package crazystring

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
		{"no w", "hi", "hi"},
		{"one letter", "w", "vv"},
		{"two ws", "wow", "vvovv"},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			if result := NoWs(tc.input); result != tc.output {
				t.Fatalf("expected %q to turn into %q but got %q instead", tc.input, tc.output, result)
			}
		})
	}
}
