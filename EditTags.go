package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func editTags(text []string) []string {
	newText := text
	for i, word := range newText {
		if strings.Contains(word, "(hex)") {
			if checkTag(word, "(hex)") {
				newText[i-1] = hexToDecimal(newText[i-1])
				newText[i] = strings.ReplaceAll(word, "(hex)", "")
				if newText[i] == "" {
					newText = append(newText[:i], newText[i+1:]...)
				}
			}
		} else if strings.Contains(word, "(bin)") {
			if checkTag(word, "(bin)") {
				newText[i-1] = binToDecimal(newText[i-1])
				newText[i] = strings.ReplaceAll(word, "(bin)", "")
				if newText[i] == "" {
					newText = append(newText[:i], newText[i+1:]...)
				}
			}
		} else if strings.Contains(word, "(up)") {
			if checkTag(word, "(up)") {
				newText[i-1] = strings.ToUpper(newText[i-1])
				newText[i] = strings.ReplaceAll(word, "(up)", "")
				if newText[i] == "" {
					newText = append(newText[:i], newText[i+1:]...)
				}
			}
		} else if strings.Contains(word, "(low)") {
			if checkTag(word, "(low)") {
				newText[i-1] = strings.ToLower(newText[i-1])
				newText[i] = strings.ReplaceAll(word, "(low)", "")
				if newText[i] == "" {
					newText = append(newText[:i], newText[i+1:]...)
				}
			}
		} else if strings.Contains(word, "(cap)") {
			if checkTag(word, "(cap)") {
				newText[i-1] = Capitalize(newText[i-1])
				newText[i] = strings.ReplaceAll(word, "(cap)", "")
				if newText[i] == "" {
					newText = append(newText[:i], newText[i+1:]...)
				}
			}
		} else if word == "(up," {
			if checkTag(word, "(up,") {
				repeat := cutSuffix(word, newText[i+1], i)
				if repeat >= 0 {
					for y := repeat; y > 0; y-- {
						newText[i-y] = strings.ToUpper(newText[i-y])
					}
				}
				newText = removeTags(newText, i)
			} else {
				fmt.Printf("Please edit your space %s\n", word)
			}
		} else if word == "(low," {
			if checkTag(word, "(low,") {
				repeat := cutSuffix(word, newText[i+1], i)
				if repeat > 0 {
					for y := repeat; y > 0; y-- {
						newText[i-y] = strings.ToLower(newText[i-y])
					}
					newText = removeTags(newText, i)
				}
			} else {
				fmt.Printf("Please edit your space %s\n", word)
			}
		} else if word == "(cap," {
			if checkTag(word, "(cap,") {
				repeat := cutSuffix(word, newText[i+1], i)
				if repeat > 0 {
					for y := repeat; y > 0; y-- {
						newText[i-y] = Capitalize(newText[i-y])
					}
					newText = removeTags(newText, i)
				}
			} else {
				fmt.Printf("Please edit your space %s\n", word)
			}
		}
	}
	for i, word := range newText {
		if word == "a" {
			newText[i] = convertA(newText[i+1])
		}
	}
	return newText
}
func cutSuffix(tag, s string, i int) int {
	suffix := "[)](.*)"
	str := regexp.MustCompile(suffix).ReplaceAllString(s, "")
	num, _ := strconv.Atoi(str)
	if num > i {
		num = i
	} else {
		fmt.Printf("Invalid Input : tag %s %s have been removed\n", tag, s)
	}
	return num
}

func cutPreffix(s string) string {
	suffix := "(.*)[)]"
	str := regexp.MustCompile(suffix).ReplaceAllString(s, "")
	return str
}

func removeTags(s []string, i int) []string {
	newText := s
	newText[i+1] = cutPreffix(newText[i+1])
	if newText[i+1] == "" {
		newText = append(newText[:i], newText[i+2:]...)
	} else {
		newText = append(newText[:i], newText[i+1:]...)
	}
	return newText
}
func checkTag(text, tag string) bool {
	patt := "^[" + tag + "].*"
	re := regexp.MustCompile(patt)
	match := re.MatchString(text)
	return match
}
