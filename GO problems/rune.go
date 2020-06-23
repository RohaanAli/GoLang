package main

import (
	"fmt"
)

func main() {
	str := "Dollars"

	chrSlice := []rune(str)

	chrSlice[0] = 'Y'
	chrSlice = append(chrSlice, '$')

	newchr := []rune{}

	for i := len(chrSlice) - 1; i >= 0; i-- {
		newchr = append(newchr, chrSlice[i])
	}

	newStr := string(newchr)

	fmt.Println(newStr)

}
