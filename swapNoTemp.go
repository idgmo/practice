package main

import (
	"fmt"
)

func main() {
	a, b := 15, 10
	fmt.Println(swap(a, b))
}

func swap(a, b int) []int {
	b, a = a, b
	return []int{a, b}
}
