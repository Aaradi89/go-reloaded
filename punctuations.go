package main

import (
	"regexp"
)

func punctuations(text string) string {
	re := regexp.MustCompile(`\s*([.,!?;:]+)\s*`)
	fixSpaces := re.ReplaceAll([]byte(text), []byte(`$1`))
	result := re.ReplaceAll([]byte(fixSpaces), []byte(`$1 `)) //add space after the .,!?;:
	return string(result)
}

func fixEndingSpace(text string) string {
	re := regexp.MustCompile(`([.,!?;:'"])\s$`)
	result := re.ReplaceAll([]byte(text), []byte(`$1`))
	re1 := regexp.MustCompile(`^ ('\S.*\S')`)
	result = re1.ReplaceAll([]byte(result), []byte(`$1`))
	re2 := regexp.MustCompile(`^ ("\S.*\S")`)
	result = re2.ReplaceAll([]byte(result), []byte(`$1`))
	return string(result)
}
