/*
commonCharacters: a function that returns a string containing only
the characters found in two input strings
To run: go build commonCharacters.go
		./commonCharacters -Istr1="string1" -Istr2="string2"
*/
package main

import (
	"flag"
	"fmt"
)

var Istr1 = flag.String("Istr1", "abcdefg", "input str1")
var Istr2 = flag.String("Istr2", "afcbedf", "input str2")

func main() {
	flag.Parse()
	fmt.Printf("string1: %s, string2 %s\n", *Istr1, *Istr2)
	fmt.Printf("characters that are common to both strings are: %s\n", commonCharacters(*Istr1, *Istr2))
}

func commonCharacters(s1, s2 string) string {
	ans := []rune{}
	hashmap := map[rune]int{}
	for _, r := range s1 {
		hashmap[r]++
	}
	for _, r := range s2 {
		if hashmap[r] > 0 {
			hashmap[r]--
			ans = append(ans, r)
		}
	}

	return string(ans)
}
