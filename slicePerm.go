package main

import (
	"fmt"
)

// Perm calls f with each permutation of a
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// perm will permute the values at index i to len(a)-1
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	var letters = []rune("abc")
	var lettersString = string(letters[:])
	var sentence = fmt.Sprintf("Permutating %s:", lettersString)
	fmt.Println(sentence)
	Perm(letters, func(b []rune) {
		fmt.Println(string(b))
	})
}
