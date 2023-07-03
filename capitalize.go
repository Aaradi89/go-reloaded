package main

func Capitalize(s string) string {
	runes := []rune(s)

	first := true

	for i := range runes {
		if first {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] -= 32
			}
			first = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
	}
	return string(runes)
}
