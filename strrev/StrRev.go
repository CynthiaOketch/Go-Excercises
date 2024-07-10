package main

import "fmt"

//	func StrRev(s string) string {
//		runes := []rune(s)
//		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//			runes[i], runes[j] = runes[j], runes[i]
//		}
//		return string(runes)
//	}
func StrRev(s string) string {
	var str2 string
	for i := range s {
		str2 += string(s[len(s)-i-1])
	}
	return str2
}

func main() {
	word := "Hello"
	fmt.Println(StrRev(word))
}
