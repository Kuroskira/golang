package main

import (
	"fmt"
)

func main() {
	v := 3
	v1 := 4
	stack := []int{}
	stack = append(stack, v)
	stack = append(stack, v1)
	fmt.Println(stack)
	stack = pop(stack)
	fmt.Println(stack)

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(s, 2))
}

func pop(stack []int) []int {
	return stack[:len(stack)-1]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
