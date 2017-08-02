package crazystring

import (
	"strings"
)

// NoWs replaces every w into two v's
func NoWs(s string) string {
	return strings.Replace(s, "w", "vv", 1)
}

// ReverseString reverse the characters of a given string
func ReverseString(s string) string {
	r := []rune(s)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}

	return string(r)
}
