package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	if len(input) == 0 {
		return
	}
	var wordStart, wordEnd int
	words := make([]string, 0)

	for i := 0; i < len(input); i++ {
		if input[i] != ' ' {
			wordStart = i
			for i < len(input) && input[i] != ' ' {
				i++
			}
			wordEnd = i
			words = append(words, input[wordStart:wordEnd])
		}
	}

	for i := len(words) - 1; i >= 0; i-- {
		if i < len(words)-1 {
			z01.PrintRune(' ')
		}
		for _, ch := range words[i] {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
