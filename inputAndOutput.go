package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput() ([]string, bool) {
	inputFile := os.Args[1]
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Input file cannot be found")
		return nil, false
	}
	defer file.Close()
	// content, _ := io.ReadAll(file)
	// return string(content), true
	scan := bufio.NewScanner(file)
	var text []string
	for scan.Scan() {
		text = append(text, scan.Text())
	}
	return text, true
}

func sendOutput(text []string) {
	outputFile := os.Args[2]
	file, _ := os.Create(outputFile)
	defer file.Close()
	writer := bufio.NewWriter(file)
	//writer.WriteString(text)
	for i, line := range text {
		if i < len(text)-1 {
			fmt.Fprintln(writer, line)
		} else {
			fmt.Fprint(writer, line)
		}
	}
	writer.Flush()
}
