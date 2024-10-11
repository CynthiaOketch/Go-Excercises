package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	for _, arg := range os.Args[1:] {
		for _, char := range processString(arg) {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
func processString(s string) string {
	var result []rune
	n := len(s)

	for i, r := range s {
		if r == ' ' {
			result = append(result, r)
		} else if i == n-1 || s[i+1] == ' ' {
			result = append(result, ToUpper(r))
		} else {
			result = append(result, ToLower(r))
		}
	}

	return string(result)
}
func ToUpper(n rune) rune {
	if n >= 'a' && n <= 'z' {
		return (n - 32)
	}
	return n
}
func ToLower(n rune) rune {
	if n >= 'A' && n <= 'Z' {
		return (n + 32)
	}
	return n
}
