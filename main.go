package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) == 2 {
		inputText, correct := getInput()
		if correct {
			// textArry := SplitInput(text)
			// text = strings.Join(textArry, " ")
			var outputText []string
			for _, line := range inputText {
				text := line
				text = fixSpaces(text)
				text = hex(text)
				text = bin(text)
				text = up(text)
				text = low(text)
				text = cap(text)
				text = specialUp(text)
				text = specialLow(text)
				text = specialCap(text)
				text = punctuations(text)
				text = fixSingleQuotations(text)
				text = fixDoubleQuotations(text)
				text = convertA(text)
				outputText = append(outputText, text)
			}
			sendOutput(outputText)
			fmt.Println("Completed")
		}
	} else if len(os.Args[1:]) < 2 {
		fmt.Println("No Input / Output file name")
	} else {
		fmt.Println("Too many arguments")
	}
}
