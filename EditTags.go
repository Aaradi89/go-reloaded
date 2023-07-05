package main

import (
	"regexp"
	"strconv"
	"strings"
)

func editTags(text []string) []string {
	newText := text
	for i, word := range newText {
		if strings.Contains(word, "(hex)") {
			newText[i-1] = hexToDecimal(newText[i-1])
			newText[i] = strings.ReplaceAll(word, "(hex)", "")
			if newText[i] == "" {
				newText = append(newText[:i], newText[i+1:]...)
			}
		} else if strings.Contains(word, "(bin)") {
			newText[i-1] = binToDecimal(newText[i-1])
			newText[i] = strings.ReplaceAll(word, "(bin)", "")
			if newText[i] == "" {
				newText = append(newText[:i], newText[i+1:]...)
			}
		} else if strings.Contains(word, "(up)") {
			newText[i-1] = strings.ToUpper(newText[i-1])
			newText[i] = strings.ReplaceAll(word, "(up)", "")
			if newText[i] == "" {
				newText = append(newText[:i], newText[i+1:]...)
			}
		} else if strings.Contains(word, "(low)") {
			newText[i-1] = strings.ToLower(newText[i-1])
			newText[i] = strings.ReplaceAll(word, "(low)", "")
			if newText[i] == "" {
				newText = append(newText[:i], newText[i+1:]...)
			}
		} else if strings.Contains(word, "(cap)") {
			newText[i-1] = Capitalize(newText[i-1])
			newText[i] = strings.ReplaceAll(word, "(cap)", "")
			if newText[i] == "" {
				newText = append(newText[:i], newText[i+1:]...)
			}
		} else if word == "(up," {
			repeat := cutSuffix(newText[i+1], i)
			if repeat > 0 {
				for y := repeat; y > 0; y-- {
					newText[i-y] = strings.ToUpper(newText[i-y])
				}
				newText = removeTags(newText, i)
			}
		} else if word == "(low," {
			repeat := cutSuffix(newText[i+1], i)
			if repeat > 0 {
				for y := repeat; y > 0; y-- {
					newText[i-y] = strings.ToLower(newText[i-y])
				}
				newText = removeTags(newText, i)
			}
		} else if word == "(cap," {
			repeat := cutSuffix(newText[i+1], i)
			if repeat > 0 {
				for y := repeat; y > 0; y-- {
					newText[i-y] = Capitalize(newText[i-y])
				}
				newText = removeTags(newText, i)
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
func cutSuffix(s string, i int) int {
	suffix := "[)](.*)"
	str := regexp.MustCompile(suffix).ReplaceAllString(s, "")
	num, _ := strconv.Atoi(str)
	if num > i {
		num = i
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
