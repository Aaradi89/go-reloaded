package main

import "fmt"

func fixQuotations(text string) string {
	for i, letter := range text {
		firstSingle := true
		//firstDouble := true
		if letter == '\'' {
			if firstSingle {
				if text[i+1] == ' ' {
					letter = ' '
					fmt.Printf("type %T", text[i+1])
				}
			}
		}
	}
	return text
}
