package main

func convertA(word string) string {
	firstLetter := rune(word[0])
	vowels := "aeoiuhAEOIUH"
	isVowel := false
	for _, vowel := range vowels {
		if firstLetter == vowel {
			isVowel = true
			break
		}
	}
	if isVowel {
		return "an"
	} else {
		return "a"
	}
}
