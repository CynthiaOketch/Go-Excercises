// inter
// Instructions

// Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

// The display will be followed by a newline ('\n').

// If the number of arguments is different from 2, the program displays nothing.

// Usage

// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
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
