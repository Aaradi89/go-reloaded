package main

import "strconv"

func hexToDecimal(s string) string {
	num, _ := strconv.ParseInt(s, 16, 64)
	result := strconv.Itoa(int(num))
	return result
}

func binToDecimal(s string) string {
	num, _ := strconv.ParseInt(s, 2, 64)
	result := strconv.Itoa(int(num))
	return result
}
