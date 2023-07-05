package main

import "fmt"

func fixSingleQuotations(text string) string {
	newTest := ""
	firstSingle := true
	for letter := 0; letter < len(text);letter++ {

		if text[letter] == '\'' && firstSingle && letter == len(text)-2 {
				if text[letter+1] == ' ' {
					newTest += " '"
					letter++
					firstSingle = false
				}
		}else if text[letter] == ' ' && !firstSingle && letter == len(text)-2 {
			if text[letter+1] == '\''{
				newTest += "' "
				letter++
				firstSingle = true
			}
		}else{
			newTest += string(text[letter])
		}
	}
	return newTest
}

func fixDoubleQuotations(text string) string {
	newTest := ""
	firstSingle := true
	for letter := 0; letter < len(text);letter++ {

		if text[letter] == '"' && firstSingle && letter == len(text)-2 {
				if text[letter+1] == ' ' {
					newTest += " \""
					letter++
					firstSingle = false
				}
		}else if text[letter] == ' ' && !firstSingle && letter == len(text)-2 {
			fmt.Println("1")
			if text[letter+1] == '"'{
				fmt.Println("2")
				newTest += "\" "
				letter++
				firstSingle = true
			}
		}else{
			newTest += string(text[letter])
		}
	}
	return newTest
}

