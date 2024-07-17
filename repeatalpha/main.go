// repeatalpha
// Instructions

// Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...
// Expected Function

// func RepeatAlpha(s string) string {
// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.RepeatAlpha("abc"))
// 	fmt.Println(piscine.RepeatAlpha("Choumi."))
// 	fmt.Println(piscine.RepeatAlpha(""))
// 	fmt.Println(piscine.RepeatAlpha("abacadaba 01!"))
// }

// And its output:

// $ go run . | cat -e
// abbccc$
// CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
// $
// abbacccaddddabba 01!$
// $
package main

import "fmt"

func RepeatAlpha(s string) string {
	var count int
	repAlpha := ""

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			count = int(c) - 'a' + 1
			for i := 0; i < count; i++ {
				repAlpha += string(c)
			}
		} else if c >= 'A' && c <= 'Z' {
			count = int(c) - 'A' + 1
			for i := 0; i < count; i++ {
				repAlpha += string(c)
			}
		} else {
			repAlpha += string(c)
		}
	}
	return repAlpha
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
