package main

import (
	"fmt"
	"strconv"
)

func HexToDec(Hex string) string {
	number, err := strconv.ParseInt(Hex, 16, 64)
	if err != nil {
		fmt.Println("Error converting to Dec", err)
	}
	return fmt.Sprint(number)
}

/*func main() {
	word := "AA"
	fmt.Println(HexToDec(word))
}
*/