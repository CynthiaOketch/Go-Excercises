// piglatin
// Instructions

// Write a program that transforms a string passed as argument in its Pig Latin version.

// The rules used by Pig Latin are as follows:

//     If a word begins with a vowel, just add "ay" to the end.
//     If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
//     Only the latin vowels will be considered as vowel(aeiou).

// If the word has no vowels, the program should print "No vowels".

// If the number of arguments is different from one, the program prints nothing.
// Usage

// $ go run .
// $ go run . pig | cat -e
// igpay$
// $ go run . Is | cat -e
// Isay$
// $ go run . crunch | cat -e
// unchcray$
// $ go run . crnch | cat -e
// No vowels$
// $ go run . something else | cat -e
// $

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
	return s == 'a' || s == 'e' || s == 'o' || s == 'i' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
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
		result := ("No Vowels")
		return result
	}
	if checkVowel(rune(s[0])) {
		result = s + "ay"
		return result
	}
	for i, ch := range s {
		if i > 0 && checkVowel(ch) {
			v := s[i:]
			s := s[:i]
			result = v + s + "ay"
		}
	}
	return result
}
