package main

import (
	"regexp"
	"strings"
)

func cap(text string) string {
	re := regexp.MustCompile(`(\S+) *\( *c *a *p *\)`)
	result := re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		getWord := string(re.ReplaceAll(b, []byte(`$1`)))
		result := Capitalize(getWord)
		return []byte(result)
	})
	return (string(result))
}

func Capitalize(s string) string {
	runes := []rune(s)

	first := true

	for i := range runes {
		if first {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] -= 32
			}
			first = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
	}
	return string(runes)
}

func up(text string) string {
	re := regexp.MustCompile(`(\S+) *\( *u *p *\)`)
	result := re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		getWord := string(re.ReplaceAll(b, []byte(`$1`)))
		result := strings.ToUpper(getWord)
		return []byte(result)
	})
	return (string(result))
}

func low(text string) string {
	re := regexp.MustCompile(`(\S+) *\( *l *o *w *\)`)
	result := re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		getWord := string(re.ReplaceAll(b, []byte(`$1`)))
		result := strings.ToLower(getWord)
		return []byte(result)
	})
	return (string(result))
}
