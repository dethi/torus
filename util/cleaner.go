package util

import (
	"regexp"
	"strings"
)

const cleanName = `(?i)((\[ *)?[a-z]+.cpasbien.[a-z]+( *\])?)|(web(-?dl)?)|(xvid)`

var regexClean = regexp.MustCompile(cleanName)

func CleanName(filename string) string {
	s, ext := SplitFilename(filename)
	s = regexClean.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ".", "-", -1)
	s = strings.Replace(s, "[", "-", -1)
	s = strings.Replace(s, "]", "-", -1)
	s = strings.Replace(s, " ", "-", -1)
	s = strings.Trim(s, "-")
	s = strings.Title(s)

	var last rune = -1
	s = strings.Map(func(c rune) rune {
		if c == '-' && last == '-' {
			return -1
		}
		last = c
		return c
	}, s)

	return s + ext
}
