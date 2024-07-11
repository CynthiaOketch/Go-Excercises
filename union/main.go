package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]
	union := make(map[rune]bool)
	unionstr := ""
	for _, c := range args[0] + args[1] {
		if !union[c] {
			unionstr += string(c)
			union[c] = true
		}
	}
	for _, c := range unionstr {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
