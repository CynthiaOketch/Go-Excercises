package main

import (
	"fmt"
	"strconv"
)

func BinToDec(Bin string) string {
	number, err := strconv.ParseInt(Bin, 2, 64)
	if err != nil {
		fmt.Println("Error converting to Dec:", err)
	}
	return fmt.Sprint(number)
}

// func main() {
// 	word := "1010111111"
// 	fmt.Println(BinToDec(word))
// }
