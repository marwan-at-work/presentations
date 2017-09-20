package crazystring

import (
	"strings"
)

// NoWs replaces every w into two v's
// NoWs(wonder) => vvonder
// NoWs(wow) => vvovv
func NoWs(s string) string {
	return strings.Replace(s, "w", "vv", 1)
}
