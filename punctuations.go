package main

func punctuations(text []string) []string {
	for i := 1; i < len(text); i++ {
		if text[i] != "" {
			if text[i][0] == '.' || text[i][0] == ',' || text[i][0] == '!' || text[i][0] == '?' || text[i][0] == ';' {
				text[i-1] += string(text[i][0])
				if len(text[i]) > 1 {
					text[i] = text[i][1:]
					i--
				} else {
					text = append(text[:i], text[i+1:]...)
				}
			}
		}
	}
	return text
}
