package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 'range' produces a pair of values: key, value.
	for idx, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println(idx, arg)
	}
	fmt.Println("Final string: ", s)

}
