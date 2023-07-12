package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func specialLow(text string) string {
	re := regexp.MustCompile(`(.*?)(\( *l *o *w *, *(\d+) *\))`)
	result := re.ReplaceAllStringFunc(text, func(s string) string {
		digit := re.ReplaceAllString(s, `$3`)
		words := re.ReplaceAllString(s, `$1`)
		tag := re.ReplaceAllString(s, `$2`)
		repeat, err := strconv.Atoi(digit)
		if err != nil {
			fmt.Printf("Note : \"%s\" is not Number. Please check %s\n", digit, tag)
			return tag
		} else {
			newRe := regexp.MustCompile(`\S+`)
			wordArry := newRe.FindAllString(words, -1)
			wordArryLen := len(wordArry) - 1
			if len(wordArry) < repeat {
				repeat = len(wordArry)
			}
			for i := wordArryLen; i > wordArryLen-repeat; i-- {
				if i >= 0 {
					wordArry[i] = strings.ToLower(wordArry[i])
				}
			}
			words = strings.Join(wordArry, " ")
		}

		return words
	})
	return result
}

func specialUp(text string) string {
	re := regexp.MustCompile(`(.*?)(\( *u *p *, *(\d+) *\))`)
	result := re.ReplaceAllStringFunc(text, func(s string) string {
		digit := re.ReplaceAllString(s, `$3`)
		words := re.ReplaceAllString(s, `$1`)
		tag := re.ReplaceAllString(s, `$2`)
		repeat, err := strconv.Atoi(digit)
		if err != nil {
			fmt.Printf("Note : \"%s\" is not Number. Please check %s\n", digit, tag)
			return tag
		} else {
			newRe := regexp.MustCompile(`\S+`)
			wordArry := newRe.FindAllString(words, -1)
			wordArryLen := len(wordArry) - 1
			if len(wordArry) < repeat {
				repeat = len(wordArry)
			}
			for i := wordArryLen; i > wordArryLen-repeat; i-- {
				if i >= 0 {
					wordArry[i] = strings.ToUpper(wordArry[i])
				}
			}
			words = strings.Join(wordArry, " ")
		}

		return words
	})
	return result
}

func specialCap(text string) string {
	re := regexp.MustCompile(`(.*?)(\( *c *a *p *, *(\d+) *\))`)
	result := re.ReplaceAllStringFunc(text, func(s string) string {
		digit := re.ReplaceAllString(s, `$3`)
		words := re.ReplaceAllString(s, `$1`)
		tag := re.ReplaceAllString(s, `$2`)
		repeat, err := strconv.Atoi(digit)
		if err != nil {
			fmt.Printf("Note : \"%s\" is not Number. Please check %s\n", digit, tag)
			return tag
		} else {
			newRe := regexp.MustCompile(`\S+`)
			wordArry := newRe.FindAllString(words, -1)
			wordArryLen := len(wordArry) - 1
			if len(wordArry) < repeat {
				repeat = len(wordArry)
			}
			for i := wordArryLen; i > wordArryLen-repeat; i-- {
				if i >= 0 {
					wordArry[i] = Capitalize(wordArry[i])
				}
			}
			words = strings.Join(wordArry, " ")
		}

		return words
	})
	return result
}
