package main

func ToLower(str string) string {
	var newstr string
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			char = char + 32
			newstr += string(char)
		} else {
			newstr += string(char)
		}
	}
	return newstr
}

// func main() {
// 	word := "HELLo"
// 	fmt.Println((ToLower(word)))
// }
