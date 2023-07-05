package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 2 {
		text := getInput()
		textArry := SplitInput(text)
		textArry = editTags(textArry)
		textArry = punctuations(textArry)
		text = strings.Join(textArry, " ")
		sendOutput(text)
		fmt.Println(text)

	} else if len(os.Args[1:]) < 2 {
		fmt.Println("No Input / Output file name")
	} else {
		fmt.Println("Too many arguments")
	}
}