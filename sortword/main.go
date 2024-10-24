// sortwordarr
// Instructions

// Write a function SortWordArr that sorts by ascii (in ascending order) a string slice.
// Expected function

// func SortWordArr(a []string) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	piscine.SortWordArr(result)

// 	fmt.Println(result)
// }

// And its output :

// $ go run .
// [1 2 3 A B C a b c]
// $
package main

import "fmt"

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}