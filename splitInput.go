package main

func SplitInput(s string) []string {
	str := ""
	var result []string
	for _, runes := range s {
		if runes == '\n' || runes == '\t' || runes == ' ' {
			if str != "" {
				result = append(result, str)
				str = ""
			}
		} else {
			str += string(runes)
		}
	}
	if str != "" {
		result = append(result, str)
	}
	return result
}
