package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	b := [5]int{0, 1, 2, 3, 4}
	reverse1(&b)
	fmt.Println(b)

	dups := []string{"hello", "ciao", "ciao"}
	dups = removeDuplicates(dups)
	fmt.Printf("%q\n", dups)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse1(s *[5]int) {
	for i := 0; i < len(s)/2; i++ {
		end := len(s) - i - 1
		s[i], s[end] = s[end], s[i]
	}
}

// Fucntion to compare to slices, we gotta make it ouself unless []byte
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func removeDuplicates(xs []string) []string {
	current := 0
	for _, v := range xs {
		if v == xs[current] {
			continue
		}
		current++
		xs[current] = v
	}
	return xs[:current+1]
}
