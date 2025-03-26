package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	if len(a) == 0 || len(nbrs) == 0 {
		return nil
	}

	start, end := 0, len(a)

	if len(nbrs) > 0 {
		start = nbrs[0]
		if start < 0 {
			start = len(a) + start
		}
	}

	if len(nbrs) > 1 {
		end = nbrs[1]
		if end < 0 {
			end = len(a) + end
		}
	}

	if start < 0 || start >= len(a) || end < 0 || end > len(a) || start >= end {
		return nil
	}

	return a[start:end]
}
