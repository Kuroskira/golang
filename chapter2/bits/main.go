package main

import (
	"fmt"
	"gobook/golang/chapter2/popcount"
)

func main() {
	var input uint64 = 28
	n := popcount.PopCount(input)
	fmt.Println(n)
}

// 101 = 5
// 1100 = 12
// 11100 = 28
