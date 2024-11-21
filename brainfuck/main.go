package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	memory := make([]byte, 2048)
	brackets := make(map[int]int)
	stack := []int{}
	pointer := 0

	for i, char := range args {
		if char == '[' {
			stack = append(stack, i)
		}
		if char == ']' {
			opening := stack[len(stack)-1]
			brackets[opening] = i
			brackets[i] = opening
			stack = stack[:len(stack)-1]
		}
	}
	for i := 0; i < len(args); i++ {
		if args[i] == '>' {
			pointer++
		}
		if args[i] == '<' {
			pointer--
		}
		if args[i] == '+' {
			memory[pointer]++
		}
		if args[i] == '-' {
			memory[pointer]--
		}
		if args[i] == '.' {
			z01.PrintRune(rune(memory[pointer]))
		}
		if args[i] == '[' && memory[pointer] == 0 {
			i = brackets[i]
		}
		if args[i] == ']' && memory[pointer] != 0 {
			i = brackets[i]
		}
	}
}
