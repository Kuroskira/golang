package main

import (
	"fmt"
	"gobook/chapter2/conv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {

		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := conv.Fahrenheit(t)
		c := conv.Celsius(t)
		k := conv.Kelvin(t)
		m := conv.Meter(t)
		ft := conv.Feet(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			f, conv.FToC(f), c, conv.CToF(c),
			k, conv.KToC(k))

		fmt.Printf("%s = %s, %s = %s\n",
			m, conv.MToFt(m), ft, conv.FtToM(ft))
	}
}
