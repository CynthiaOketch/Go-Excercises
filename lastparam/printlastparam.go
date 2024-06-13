package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i, j := range arg {
		if i == len(arg)-1 {
			for _, c := range j {
				z01.PrintRune(c)
			}
		}
	}
	z01.PrintRune('\n')
}
