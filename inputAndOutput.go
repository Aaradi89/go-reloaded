package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getInput() (string, bool) {
	inputFile := os.Args[1]
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Input file cannot be found")
		return "", false
	}
	defer file.Close()
	content, _ := io.ReadAll(file)
	return string(content), true
}

func sendOutput(text string) {
	outputFile := os.Args[2]
	file, _ := os.Create(outputFile)
	writer := bufio.NewWriter(file)
	writer.WriteString(text)
	writer.Flush()
}
