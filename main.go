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
			textArry = editTags(textArry)
			textArry = punctuations(textArry)
			text = strings.Join(textArry, " ")
			text = fixSingleQuotations(text)
			text = fixDoubleQuotations(text)
			sendOutput(text)
			fmt.Println("Completed")
		}
		//tester()
	} else if len(os.Args[1:]) < 2 {
		fmt.Println("No Input / Output file name")
	} else {
		fmt.Println("Too many arguments")
	}
}

// func tester() {
// 	hexTest := hexToDecimal("p")
// 	binaryTest := binToDecimal("0a")
// 	capitalizeTest := Capitalize("this")
// 	fmt.Printf("Hex Test : %s\n", hexTest)
// 	fmt.Printf("Binary Test : %s\n", binaryTest)
// 	fmt.Printf("Capitalize Test : %s\n", capitalizeTest)
// }
