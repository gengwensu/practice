/*
Is Unique: Implement an algorithm to determine if a string has all unique characters
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
	m := map[byte]int{} //create an empty map
	for i := 0; i < len(str); i++ {
		m[str[i]]++
		if m[str[i]] > 1 {
			return false
		}
	}
	return true
}
