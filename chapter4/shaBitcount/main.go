package main

import (
	"fmt"
	"golang/chapter2/popcount"
)

func main() {
	var x uint64 = 34

	bits := popcount.PopCount(x)
	fmt.Println(bits)
}
