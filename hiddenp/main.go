package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1, str2 := os.Args[1], os.Args[2]
	index1, index2 := 0, 0
	for index1 < len(str1) && index2 < len(str2) {
		if str1[index1] == str2[index2] {
			index1++
		}
		index2++
	}
	if index1 == len(str1) {
		z01.PrintRune('1')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}
