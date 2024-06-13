package main

func Max(slice []int) int {
	var Great int
	for _, c := range slice {
		if c > Great {
			Great = c
		}
	}
	return Great
}

// func main() {
// 	numbers := []int64{1, 4, 3, 789898, 2}
// 	fmt.Println(Max(numbers))
// }
