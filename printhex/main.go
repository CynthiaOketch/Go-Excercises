// printhex
// Instructions

// Write a program that takes a positive (or zero) number expressed in base 10, and that displays it in base 16 (with lowercase letters) followed by a newline ('\n').

//     If the number of arguments is different from 1, the program displays nothing.
//     Error cases have to be handled as shown in the example below.

// Usage

// $ go run . 10
// a
// $ go run . 255
// ff
// $ go run . 5156454
// 4eae66
// $ go run .
// $ go run . "123 132 1" | cat -e
// ERROR$
// $
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		res := printHex(n, "0123456789abcdef")
		fmt.Println(res)
	}
}

func printHex(num int, base string) string {
	result := ""

	for num > 0 {
		rem := num % len(base)
		result = string(base[rem]) + result
		num /= len(base)
	}
	return result
}
