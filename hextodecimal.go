package main

import "strconv"

func hexToDecimal(s string) int {
	num, _ := strconv.ParseInt(s, 16, 64)
	result := int(num)
	return result
}

func binToDecimal(s string) int {
	num, _ := strconv.ParseInt(s, 2, 64)
	result := int(num)
	return result
}
