package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}
	args := os.Args[1]
	splitted := split(args, " ")
	joined := join(splitted)
	for _, char := range joined {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
func split(str, sep string) []string {
	var result []string
	start := 0
	for i := 0; i+len(sep) <= len(str); i++ {
		if str[i:i+len(sep)] == sep {
			if start < i {
				result = append(result, str[start:i])
			}
			start = i + len(sep)
		}
	}
	if start < len(str) {
		result = append(result, str[start:])
	}
	return result
}
func join(slice []string) string {
	final := ""
	for i, word := range slice {
		final += word
		if i < len(slice)-1 {
			final += "   "
		}
	}
	return final
}
