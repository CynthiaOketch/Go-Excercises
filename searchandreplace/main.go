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
	if len(os.Args) != 3 {
		return
	}

	args1, args2, seen, exist := os.Args[1], os.Args[2], make(map[rune]bool), make(map[rune]bool)

	for _, char := range args2 {
		exist[char] = true
	}

	for _, char := range args1 {
		if exist[char] && !seen[char] {
			seen[char] = true
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
