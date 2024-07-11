// searchreplace
// Instructions

// Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

//     If the number of arguments is different from 3, the program displays nothing.

//     If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').

// Usage

// $ go run . "hella there" "a" "o"
// hello there
// $ go run . "hallo thara" "a" "e"
// hello there
// $ go run . "abcd" "z" "l"
// abcd
// $ go run . "something" "a" "o" "b" "c"
// $
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str1 := os.Args[1]
	str2 := os.Args[2]
	str3 := os.Args[3]
	str := ""
	for _, c := range str1 {
		if string(c) == str2 {
			str += str3
		} else {
			str += string(c)
		}
	}
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
