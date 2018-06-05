package main

import "fmt"

func main() {
	var a [3]int // Array of three integers
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("idx: %d value: %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("value: %d\n", v)
	}

	// Literal array initialization
	//var q [3]int = [3]int{1, 2, 3}
	q := [...]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
	fmt.Printf("%T\n", q)

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	for i, v := range symbol {
		fmt.Printf("idx: %d value: %d\n", i, v)
	}

	a1 := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a1 == b, a1 == c, b == c)
	// d := [3]int{1, 2}
	// fmt.Println(a1 == d)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
