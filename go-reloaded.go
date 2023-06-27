package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args[0:]) < 2 {
		fmt.Println("File name is missing")
	} else if len(os.Args[0:]) == 2 {
		fileName := os.Args[1]
		content, _ := ioutil.ReadFile(fileName)
		fmt.Println(string(content))
	} else {
		fmt.Println("Too many arguments")
	}
}
