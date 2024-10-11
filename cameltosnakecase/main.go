package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	iscamelcase := true

	//check if the string is valid camelcase
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z' {
				iscamelcase = false
			}
		} else if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			iscamelcase = false
		}
	}

	if !iscamelcase || (s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z') {
		return s
	}
	result := ""
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(char)
		} else {
			result += string(char)
		}
	}
	return result
}
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))       // hello_World
	fmt.Println(CamelToSnakeCase("camelCase"))        // camel_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("hey2"))             // hey2
}
