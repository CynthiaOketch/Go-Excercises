// doop
// Instructions

// Write a program that is called doop.

// The program has to be used with three arguments:

//     A value
//     An operator, one of : +, -, /, *, %
//     Another value

// In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

// The program has to handle the modulo and division operations by 0 as shown on the output examples below.
// Usage

// $ go run .
// $ go run . 1 + 1 | cat -e
// 2
// $
// $ go run . hello + 1
// $ go run . 1 p 1
// $ go run . 1 / 0 | cat -e
// No division by 0
// $
// $ go run . 1 % 0 | cat -e
// No modulo by 0
// $
// $ go run . 9223372036854775807 + 1
// $ go run . -9223372036854775809 - 3
// $ go run . 9223372036854775807 "*" 3
// $ go run . 1 "*" 1
// 1
// $ go run . 1 "*" -1
// -1
// $

package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	value1, operator, value2 := os.Args[1], os.Args[2], os.Args[3]

	result := Calculation(value1, operator, value2)

	if result != "" {
		os.Stdout.WriteString(result)
		os.Stdout.WriteString("\n")
	}
}

func Calculation(a, operator, b string) string {
	value1, ok1 := Atoi(a)
	value2, ok2 := Atoi(b)

	if !ok1 || !ok2 {
		return ""
	}
	if (value1 >= 9223372036854775807 || value1 <= -9223372036854775807) || (value2 >= 9223372036854775807 || value2 <= -9223372036854775807) {
		return ""
	}
	q := 0
	switch operator {
	case "+":
		q = value1 + value2
	case "-":
		q = value1 - value2
	case "*":
		q = value1 * value2
	case "/":
		if value2 == 0 {
			return "No division by 0"
		}
		q = value1 / value2
	case "%":
		if value2 == 0 {
			return "No modulo by 0"
		}
		q = value1 % value2
	default:
		return ""
	}
	return Itoa(q)
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	q := ""
	if n > 0 {
		digits := n % 10
		q = string(rune('0'+digits)) + q
		n /= 10
	}
	return sign + q
}

func Atoi(s string) (int, bool) {
	sign := 1
	q := 0

	for i, v := range s {
		if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else if v >= '0' && v <= '9' {
			q = q*10 + int(v-'0')
		} else {
			return 0, false
		}
	}
	return sign * q, true
}
