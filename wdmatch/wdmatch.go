package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	str1 := args[0]
	str2 := args[1]

	index1 := 0
	index2 := 0

	for index1 < len(str1) && index2 < len(str2) {
		if str1[index1] == str2[index2] {
			index1++
		}
		index2++
	}
	if index1 == len(str1) {
		fmt.Println(str1)
	}
}
