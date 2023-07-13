package main

import "regexp"

func fixSpaces(text string) string {
	re := regexp.MustCompile(` +`)
	newText := re.ReplaceAll([]byte(text), []byte(` `))
	return string(newText)
}
