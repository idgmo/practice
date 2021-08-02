package main

import (
	"fmt"
)

// min returns smaller of x or y
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// max returns larger of x or y
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(Min(5, 10))
	fmt.Println(Max(5, 10))
}
