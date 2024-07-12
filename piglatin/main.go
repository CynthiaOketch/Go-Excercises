package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	fmt.Println(pigLatin(args))
}

func checkVowel(s rune) bool {
	return s == 'a' || s == 'e' || s == 'o' || s == 'i' || s == 'u'
}

func pigLatin(s string) (result string) {
	check := false

	for _, char := range s {
		if checkVowel(char) {
			check = true
		}
		if check {
			break
		}
	}
	if !check {
		result := ("No Vowel")
		return result
	}
	if checkVowel(rune(s[0])) {
		result = s + "ay"
		return result
	}
	for i, ch := range s {
		if i > 0 && checkVowel(ch) {
			v := s[i:]
			s = s[:i]
			result = v + s + "ay"
		}
	}
	return result
}
