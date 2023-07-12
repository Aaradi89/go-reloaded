package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 2 {
		text, correct := getInput()
		if correct {
			textArry := SplitInput(text)
			text = strings.Join(textArry, " ")
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
			sendOutput(text)
			fmt.Println("Completed")
		}
	} else if len(os.Args[1:]) < 2 {
		fmt.Println("No Input / Output file name")
	} else {
		fmt.Println("Too many arguments")
	}
}
