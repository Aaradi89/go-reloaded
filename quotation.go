package main

import "strings"

func fixSingleQuotations(text string) string {
	newText := ""
	firstSingle := true
	for letter := 0; letter < len(text)-1; letter++ {
		if text[letter:letter+2] == "' " && firstSingle {
			newText += " '"
			letter++
			firstSingle = false
		} else if text[letter:letter+2] == " '" && !firstSingle {
			newText += "' "
			letter++
			firstSingle = true
		} else {
			newText += string(text[letter])
		}
	}
	// if newText[len(newText)-1] == ' ' {
	// 	newText = newText[:len(newText)-1]
	// }
	return newText
}

func fixDoubleQuotations(text string) string {
	newText := ""
	firstSingle := true
	for letter := 0; letter < len(text); letter++ {
		if text[letter] == '"' && firstSingle && letter < len(text)-1 {
			if text[letter+1] == ' ' {
				newText += " \""
				letter++
				firstSingle = false
			}
		} else if text[letter] == ' ' && !firstSingle && letter < len(text)-1 {
			if text[letter+1] == '"' {
				newText += "\" "
				letter++
				firstSingle = true
			}
		} else {
			newText += string(text[letter])
		}
	}
	// if newText[len(newText)-1] == ' ' {
	// 	newText = newText[:len(newText)-1]
	// }
	cleanArry := SplitInput(newText)
	newText = strings.Join(cleanArry, " ")
	return newText
}

// func fixSingleQuotationsTest(text string) string {
// 	newText := text
// 	re := regexp.MustCompile("' *")
// 	matcher := re.MatchString(newText)
// 	fmt.Println(matcher)
// 	return newText
// }
