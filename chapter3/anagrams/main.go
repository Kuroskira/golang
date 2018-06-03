package main

import "fmt"

func main() {
	name := "hello"
	name1 := "olleh"
	isAnagram := isAnagram(name, name1)
	fmt.Println(isAnagram)
}

func isAnagram(s1, s2 string) bool {
	countS1 := make(map[rune]int)
	for _, c := range s1 {
		countS1[c]++
	}

	countS2 := make(map[rune]int)
	for _, c := range s2 {
		countS2[c]++
	}

	for k, v := range countS1 {
		if countS2[k] != v {
			return false
		}
	}

	for k, v := range countS2 {
		if countS1[k] != v {
			return false
		}
	}

	return true
}
