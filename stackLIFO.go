package main

import (
	"fmt"
)

func main() {

	// create variable
	var stack []string

	// push data to variable
	stack = append(stack, "world!\n")
	stack = append(stack, "Hello ")
	stack = append(stack, "Well ")

	// print out the last value and then remove it
	for len(stack) > 0 {
		// Print top of slice
		n := len(stack) - 1
		fmt.Print(stack[n])

		// Pop
		stack = stack[:n]
	}
	//stack is empty at the end of the loop

	// Output: Hello world!
}
