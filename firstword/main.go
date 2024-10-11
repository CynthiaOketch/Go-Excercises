package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
func FirstWord(s string) string {
	var result string
	for _, char := range s {
		result += string(char)
		if char == ' ' {
			break
		}
	}
	return result + "\n"
}
