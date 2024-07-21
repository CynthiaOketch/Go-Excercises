package main

import "fmt"

func main() {
	result := Unique("foo", "boo")
	fmt.Printf("The string has %d unique characters.\n", result)
}

func Unique(s1, s2 string) int {
	uniquechars := make(map[rune]int) // map to store the count of the characters

	for _, char := range s1 {
		uniquechars[char]++
	}
	for _, char := range s2 {
		uniquechars[char]++
	}
	uniquecount := 0
	for _, val := range uniquechars {
		if val == 1 {
			uniquecount++
		}
	}
	return uniquecount
}
