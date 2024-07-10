package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	const rot13 = 13
	s := os.Args[1]
	for _, c := range s {
		if c >= 'a' && c <= 'z' && c+rot13 > 'z' {
			c = (c + rot13) - 26
		} else if c >= 'A' && c <= 'Z' && c+rot13 > 'Z' {
			c = (c + rot13) - 26
		} else if c >= 'a' && c <= 'z' && c+rot13 < 'z' {
			c = c + rot13
		} else if c >= 'A' && c <= 'Z' && c+rot13 < 'Z' {
			c = c + rot13
		} else {
			c = c + rot13
		}
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
