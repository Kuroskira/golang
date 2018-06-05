package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var sh384 = flag.Bool("n", false, "SHA384")
var sh512 = flag.Bool("m", false, "SHA512")

func main() {
	//hash := make([]byte, 0, 64)
	flag.Parse()
	arg := flag.Args()

	bytes := mapList(arg, toBytes)

	if !*sh384 {
		hash := sha512.Sum384(bytes[0])
		fmt.Printf("%x\n", hash)
	} else if !*sh512 {
		hash := sha512.Sum512(bytes[0])
		fmt.Printf("%x\n", hash)
	} else {
		hash := sha256.Sum256(bytes[0])
		fmt.Printf("%x\n", hash)
	}
}

func mapList(vs []string, f func(string) []byte) [][]byte {
	vsm := make([][]byte, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func toBytes(s string) []byte {
	return []byte(s)
}
