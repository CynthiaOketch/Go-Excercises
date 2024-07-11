// ispowerof2
// Instructions

// Write a program that determines if a given number is a power of 2. A number n is a power of 2 if it meets the following condition: n = 2 m where m is another integer number.

// For example, 4 = 2 2 or 16 = 2 4 are power of 2 numbers while 24 is not.

// This program must print true if the given number is a power of 2, otherwise it has to print false.

// If there is more than one or no argument, the program should print nothing.

// When there is only one argument, it will always be a positive valid int.

// Usage :

// $ go run . 2 | cat -e
// true$
// $ go run . 64
// true
// $ go run . 513
// false
// $ go run .
// $ go run . 64 1024
// $
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	numb := Atoi(os.Args[1])

	var ans string
	if isPowerOf2(numb) {
		ans = "true"
	} else {
		ans = "false"
	}
	for _, v := range ans {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func isPowerOf2(numb int) bool {
	return numb > 0 && (numb&(numb-1)) == 0
}

func Atoi(s string) int {
	var number int
	sign := 1
	for index, c := range s {
		if index == 0 && c == '-' {
			sign = -1
		} else if c == '+' && index == 0 {
			sign = 1
		} else if c >= '0' || c <= '9' {
			number = number*10 + int(c-'0')
		} else {
			return 0
		}
	}
	return sign * number
}
