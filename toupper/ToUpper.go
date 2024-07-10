package main

import "fmt"

func ToUpper(str string) string {
	var newstring string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			c = c - 32
			newstring += string(c)
		} else {
			newstring += string(c)
		}
	}
	return newstring
}

/*func main() {
	word := "Hello there ia am cynthia"
	fmt.Println(ToUpper(word))
}
