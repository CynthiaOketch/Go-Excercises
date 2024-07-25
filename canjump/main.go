package main

import "fmt"

func CanJump(nums []uint) bool {
	if int(len(nums)) == 0 {
		return false
	}
	if int(len(nums)) == 1 {
		return true
	}

	pos := 0
	for i := 0; i < len(nums); i += pos {
		pos = int(nums[i])
		if pos == 0 && i < len(nums)-1 {
			return false
		}
		if i+pos == len(nums)-1 {
			return true
		}
	}
	return false
}
func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println("Example 1:", CanJump(input1)) // true

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println("Example 2:", CanJump(input2)) // false

	input3 := []uint{0}
	fmt.Println("Example 3:", CanJump(input3)) // true

	input4 := []uint{1, 1, 1, 1, 1}
	fmt.Println("Example 4:", CanJump(input4)) // true

	input5 := []uint{2, 0, 2, 0, 1}
	fmt.Println("Example 5:", CanJump(input5)) // true

	input6 := []uint{1, 0, 1, 0}
	fmt.Println("Example 6:", CanJump(input6)) // false

	input7 := []uint{4, 3, 2, 1, 0}
	fmt.Println("Example 7:", CanJump(input7)) // true
}
