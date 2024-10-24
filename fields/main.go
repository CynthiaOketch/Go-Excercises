package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Custom implementation of strings.Fields
func fields(str string) []string {
	var result []string
	start := -1 // Track the start of each word

	// Iterate over each character in the string
	for i, r := range str {
		// Check if the character is not a whitespace
		if !unicode.IsSpace(r) {
			// If start is -1, it means we are starting a new word
			if start == -1 {
				start = i
			}
		} else {
			// If start is not -1, it means we were inside a word
			if start != -1 {
				// Append the word to the result
				result = append(result, str[start:i])
				start = -1 // Reset start since we're between words
			}
		}
	}

	// Handle the case where the string ends with a non-whitespace character
	if start != -1 {
		result = append(result, str[start:])
	}

	return result
}

func main() {
	// Test input
	s := "  Go           is   awesome     and powerful  "

	// Use the custom fields function
	words := fields(s)
	splitted := split(s, " ")
	joined := strings.Join(words, " ")
	joined2 := strings.Join(splitted, " ")

	// Output the resulting slice of words
	fmt.Println(words) // Output: ["Go", "is", "awesome", "and", "powerful"]
	fmt.Println(joined)
	fmt.Println(joined2)
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
