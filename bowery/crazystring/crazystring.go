package crazystring

import (
	"strings"
)

// NoWs replaces every w into two v's
func NoWs(s string) string {
	return strings.Replace(s, "w", "vv", 1)
}
