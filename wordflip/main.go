package main

import (
	"fmt"
)

func Split(str, sep string) []string {
	var result []string
	start := 0
	for i := 0; i+len(sep) <= len(str); i++ {
		if str[i:i+len(sep)] == sep {
			if start < i {
				result = append(result, str[start:i])
			}
			start = i + len(sep)
		}
	}
	if start < len(str) {
		result = append(result, str[start:])
	}
	return result
}

func WordFlip(str string) string {
	var final string
	res := Split(str, " ")
	for i := len(res) - 1; i >= 0; i-- {
		if i > 0{
			final += res[i] + " "
		} else if i == 0 {
			final += res[i]
		}
		
	}
	//joined := strings.Join(final, " ")
	return final
}

func main() {
	fmt.Println(WordFlip("First second last"))
	fmt.Println(WordFlip(""))
	fmt.Println(WordFlip("     "))
	fmt.Println(WordFlip(" hello  all  of  you! "))
}
