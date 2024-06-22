package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	args := os.Args[1]
	var mirror string
	for _, c := range args {
		if c >= 'A' && c <= 'Z' {
			c = ('Z' - c) + 'A'
			mirror += string(c)
		} else if c >= 'a' && c <= 'z' {
			c = ('z' - c) + 'a'
			mirror += string(c)
		} else {
			mirror += string(c)
		}
	}
	for _, c := range mirror {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
