package main

func Up(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] -= 32
		}
	}
	return string(runes)
}

func Low(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
	}
	return string(runes)
}
