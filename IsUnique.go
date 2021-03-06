/*
Is Unique: Implement an algorithm to determine if a string has all unique characters
The algorithm works with UTF-8 coded string. For ASCII, it could be even simpler.
*/
package main

import (
	"flag"
	"fmt"
)

var Istr = flag.String("Istr", "abc", "input str for IsUnique")

func main() {
	flag.Parse()
	fmt.Printf("Input string %s\n", *Istr)
	if isUnique(*Istr) {
		fmt.Println(*Istr, " is Unique!")
	} else {
		fmt.Println(*Istr, " is not unique!")
	}
}
func isUnique(str string) bool {
	m := map[rune]int{} //create an empty map
	for _, r := range str {
		m[r]++
		if m[r] > 1 {
			return false
		}
	}
	return true
}
