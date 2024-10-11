package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

// Gcd calculates the greatest common divisor of two strictly positive integers.
func Gcd(a, b uint) uint {
	// If any number is 0, return 0.
	if a == 0 || b == 0 {
		return 0
	}

	// Using the Euclidean algorithm to find GCD
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
