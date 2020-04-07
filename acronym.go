package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns the string abbreviation
func Abbreviate(s string) string {
	s = strings.ToUpper(regexp.MustCompile(`[^a-zA-Z0-9- ]`).ReplaceAllString(s, ""))
	return strings.Join(regexp.MustCompile(`\b\w`).FindAllString(s, -1), "")
}
