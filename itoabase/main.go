package main

import "fmt"



func main() {
	fmt.Println(ItoaBase(10, 2))
	fmt.Println(ItoaBase(255, 16))
	fmt.Println(ItoaBase(-42, 4))
	fmt.Println(ItoaBase(123, 10))
	fmt.Println(ItoaBase(0, 8))
	fmt.Println(ItoaBase(255, 2))
	fmt.Println(ItoaBase(-255, 16))
	fmt.Println(ItoaBase(15, 16))
	fmt.Println(ItoaBase(10, 4))
	fmt.Println(ItoaBase(255, 10))
}

func ItoaBase(num, b int) string {
	if num == 0 {
		return "0"
	}
	hexv := "0123456789ABCDEF"
	res := ""
	sign := ""
	if num < 0 {
		num *= -1
		sign = "-"
	}
	value := uint(num)
	base := uint(b)
	for value > 0 {
		dig := value % base
		res = string(hexv[dig]) + res
		value /= base
	}
	return sign + res
}
