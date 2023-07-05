package main

import (
	"bufio"
	"io"
	"os"
)

func getInput() string {
	inputFile := os.Args[1]
	file, _ := os.Open(inputFile)
	defer file.Close()
	content, _ := io.ReadAll(file)
	return string(content)
}

func sendOutput(text string) {
	outputFile := os.Args[2]
	file, _ := os.Create(outputFile)
	writer := bufio.NewWriter(file)
	writer.WriteString(text)
	writer.Flush()
}
