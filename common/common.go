package common

import "regexp"

var escape_regexp = regexp.MustCompile("[\\\"]")

func Escape(s string) string {
	return escape_regexp.ReplaceAllString(s, "\\${0}")
}
