package main

import "fmt"

func Itoa(nb int) string {
	var result string
	if nb == 0 {
		return "0"
	}
	for nb > 0 {
		result = string('0'+(nb%10)) + result
		nb /= 10
	}
	return result
}

func main() {
	number := 123
	fmt.Println(Itoa(number))
}
