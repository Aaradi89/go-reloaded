package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func getInput() ([]string, bool) {
	inputFile := os.Args[1]
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Input file cannot be found")
		return nil, false
	}
	defer file.Close()
	re := regexp.MustCompile(`.txt$`)
	if !re.MatchString(inputFile) || !re.MatchString(os.Args[2]) {
		fmt.Println("Error : Input/Output file extension should be .txt")
		return nil, false
	}
	scan := bufio.NewScanner(file)
	var text []string
	for scan.Scan() {
		text = append(text, scan.Text())
	}
	return text, true
}

func sendOutput(text []string) {
	outputFile := os.Args[2]
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i, line := range text {
		if i < len(text)-1 {
			fmt.Fprintln(writer, line)
		} else {
			fmt.Fprint(writer, line)
		}
	}
	writer.Flush()
}
