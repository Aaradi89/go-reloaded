package main

import (
	"regexp"
)

func fixSingleQuotations(text string) string {
	re := regexp.MustCompile(`\s*'\s*(.*?)\s*'\s*`)
	result := string(re.ReplaceAll([]byte(text), []byte(` '$1' `)))
	return result
}

func fixDoubleQuotations(text string) string {
	re := regexp.MustCompile(`\s*"\s*(.*?)\s*"\s*`)
	result := string(re.ReplaceAll([]byte(text), []byte(` "$1" `)))
	return result
}
