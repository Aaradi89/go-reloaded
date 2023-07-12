package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func hex(text string) string {
	re := regexp.MustCompile(`(\S+) *\( *h *e *x *\)`)
	result := re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		getHex := string(re.ReplaceAll(b, []byte(`$1`)))
		result := hexToDecimal(getHex)
		return []byte(result)
	})
	return (string(result))
}

func bin(text string) string {
	re := regexp.MustCompile(`(\S+) *\( *b *i *n *\)`)
	result := re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		getBin := string(re.ReplaceAll(b, []byte(`$1`)))
		result := hexToDecimal(getBin)
		return []byte(result)
	})
	return (string(result))
}

func hexToDecimal(s string) string {
	num, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Printf("Note : \"%s\" is not a Hex Number. it remain unchanged\n", s)
		return s + `(hex)`
	}
	result := strconv.Itoa(int(num))
	return result
}

func binToDecimal(s string) string {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Printf("Note : \"%s\" is not a Binary Number. it remain unchanged\n", s)
		return s + `(bin)`
	}
	result := strconv.Itoa(int(num))
	return result
}
