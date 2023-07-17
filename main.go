package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) == 2 {
		inputText, correct := getInput()
		if correct {
			finalCheck := true
			var outputText []string
			for i, line := range inputText {
				text := line
				text = punctuations(text)
				text = fixSingleQuotations(text)
				text = fixDoubleQuotations(text)
				text = editTags(text)
				if checkTags(text) {
					text = convertA(text)
					text = fixEndingSpace(text)
					outputText = append(outputText, text)
				} else {
					fmt.Println("Please fix the tags above in pargraph number", i+1, "and try again")
					finalCheck = false
				}
			}
			if finalCheck {
				sendOutput(outputText)
				fmt.Println("Completed")
			}
		}
	} else if len(os.Args[1:]) < 2 {
		fmt.Println("No Input / Output file name")
	} else {
		fmt.Println("Too many arguments")
	}
}
