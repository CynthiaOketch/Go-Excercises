// gcd
// Instructions

// Write a function that takes two uint representing two strictly positive integers and returns their greatest common divisor. If any of the input numbers is 0, the function should return 0.

//     In mathematics, the greatest common divisor (GCD) of two or more integers, which are not all zero, is the largest positive integer that divides each of the integers.

// Expected function

// func Gcd(a, b uint) uint {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Gcd(42, 10))
// 	fmt.Println(piscine.Gcd(42, 12))
// 	fmt.Println(piscine.Gcd(14, 77))
// 	fmt.Println(piscine.Gcd(17, 3))
// }

// And its output :

// $ go run .
// 2
// 6
// 7
// 1
// $
package main

import "fmt"

func Gcd(a, b int) int {
	gcd := 0
	if a <= 0 || b <= 0 {
		return 0
	}
	if a > b {
		for i := 1; i <= a; i++ {
			if a%i == 0 && b%i == 0 {
				gcd = i
			}
		}
	} else if b > a {
		for i := 1; i <= b; i++ {
			if a%i == 0 && b%i == 0 {
				gcd = i
			}
		}
	}
	return gcd
}

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
