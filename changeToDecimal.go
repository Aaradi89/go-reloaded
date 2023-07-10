package main

import (
	"fmt"
	"strconv"
)

func hexToDecimal(s string) string {
	num, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Printf("Note : \"%s\" is not a Hex Number. it remain unchanged\n", s)
		return s
	}
	result := strconv.Itoa(int(num))
	return result
}

func binToDecimal(s string) string {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Printf("Note : \"%s\" is not a Binary Number. it remain unchanged\n", s)
		return s
	}
	result := strconv.Itoa(int(num))
	return result
}
