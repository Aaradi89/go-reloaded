package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var input, output, outputFile string

	if len(os.Args[0:]) < 3 {
		fmt.Println("File name is missing")
	} else if len(os.Args[0:]) == 3 {
		fileName := os.Args[1]
		content, _ := ioutil.ReadFile(fileName)
		input = string(content)
		outputFile = os.Args[2]
	} else {
		fmt.Println("Too many arguments")
	}
	outputArry := SplitInput(input)
	for _, word := range outputArry {
		output += word
	}
	//might need to be fix
	result, _ := os.Create(outputFile)
	result.WriteString(output)
	//ioutil.WriteFile(outputFile, []byte(output), 0644)

	fmt.Println(hexToDecimal("1E"))
	fmt.Println(binToDecimal("1001"))
	fmt.Println(Capitalize("f**sGREHb dffd!@$#234v gyuf"))
	fmt.Println(Up("awvgrw12cda!@$vwe"))
	fmt.Println(Low("VDSAV23rfdwsE@#fWRVRWG"))
	fmt.Println(convertA("HRli"))
	fmt.Println(convertA("hiariam"))
	test := []string{"", "asd", ",vty", "serx", ";!vuvuy", "!?", "", "byvyv", "."}
	fmt.Println(punctuations(test))
	fmt.Println(fixQuotations("' graiopgjalm ' ' "))
}
