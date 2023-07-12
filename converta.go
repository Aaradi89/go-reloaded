package main

import (
	"regexp"
)

func convertA(s string) string {
	re := regexp.MustCompile(`(\s[aA])( [aeoiuhAEOIUH])`)
	result := re.ReplaceAll([]byte(s), []byte(`${1}n$2`))
	result = re.ReplaceAll([]byte(result), []byte(`${1}n$2`))
	return string(result)
}
