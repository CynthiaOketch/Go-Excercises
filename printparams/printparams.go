package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for _, char := range arg {
		for _, runes := range char {
			z01.PrintRune(runes)
		}
		z01.PrintRune('\n')
	}
}
