package main

func isAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func Capitalize(s string) string {
	words := []byte(s)
	capitalizeNext := true
	for i := 0; i < len(words); i++ {
		if isAlphanumeric(words[i]) {
			if capitalizeNext {
				words[i] = toUpper(words[i])
				capitalizeNext = false
			} else {
				words[i] = toLower(words[i])
			}
		} else {
			capitalizeNext = true
		}
	}
	return string(words)
}

func toUpper(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return char - ('a' - 'A')
	}
	return char
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

// func main() {
// 	word := "Hello! How are you? How@are@things+4you?"
// 	fmt.Println(Capitalize(word))
// }
