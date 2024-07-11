// compare
// Instructions

// Write a function that behaves like the Compare function.
// Expected function

// func Compare(a, b string) int {

// }

// Usage

// Here is a possible program to test your function :

// And its output :

// $ go run .
// 0
// -1
// 1
// $

package main

import (
	"fmt"
)

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	} else {
		return 1
	}
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
