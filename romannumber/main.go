package main

import (
	"errors"
	"fmt"
	"os"
)

func PrintRoman(nb int) (calculated, symbols string) {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	calculation := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var calc string
	var symb string

	for i := 0; i <= len(values)-1; i++ {
		for nb >= values[i] {
			nb -= values[i]
			symb += symbol[i]

			if calc != "" {
				calc += "+"
			}
			calc += calculation[i]
		}
	}
	return calc, symb
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run . value")
		return
	}
	val, err := Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	if val <= 0 || val >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	calc, symb := PrintRoman(val)
	fmt.Println(calc)
	fmt.Println(symb)
}

func Atoi(s string) (int, error) {
	var number int
	sign := 1
	for i, char := range s {
		if i == 0 && char == '+' {
			sign = 1
		} else if i == 0 && char == '-' {
			sign = -1
		} else if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		} else {
			return 0, errors.New("ERROR: cannot convert to roman digit")
		}
	}
	return number * sign, nil
}
