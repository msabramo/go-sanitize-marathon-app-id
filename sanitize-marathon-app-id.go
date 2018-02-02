package sanitizemarathonappid

import (
	"regexp"
	"strings"
)

var (
	charsToReplaceRegexp        = regexp.MustCompile(`[\._]`)
	multipleHyphensRegexp       = regexp.MustCompile(`-{2,}`)
	nonSubstitutableCharsRegexp = regexp.MustCompile(`[^0-9a-z-]`)
)

// Sanitize takes an input string and sanitizes it so it is a legal Marathon app id
func Sanitize(s string) string {
	s = strings.ToLower(s)
	s = charsToReplaceRegexp.ReplaceAllString(s, "-")
	s = nonSubstitutableCharsRegexp.ReplaceAllString(s, "")
	s = multipleHyphensRegexp.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
