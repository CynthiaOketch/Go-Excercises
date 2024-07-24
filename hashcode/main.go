package main

import (
	"fmt"
)

func HashCode(dec string) string {
	size := len(dec)
	hashed := ""

	for _, char := range dec {
		ascii := int(char) + size
		hashValue := ascii % 127

		// Ensure hash value is printable
		if hashValue < 32 {
			hashValue += 33
		}

		hashed += string(hashValue)
	}

	return hashed
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
	fmt.Println(HashCode("14 Avril 1912"))
	fmt.Println(HashCode("14 Avril@1912"))
}
