/*
Anagram: a function that reports whether two strings are anagrams of each other,
that is, they contain the same letters in a dif ferent order.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/gengwensu/mylib"
)

var Istr1 = flag.String("Istr1", "abc", "input str1 for anagrams")
var Istr2 = flag.String("Istr2", "bca", "input str2 for anagrams")

func main() {
	flag.Parse()
	fmt.Printf("string1: %s, string2 %s\n", *Istr1, *Istr2)
	fmt.Printf("Are these two strings anagram? %v\n", anagrams(*Istr1, *Istr2))
}

func anagrams(s1, s2 string) bool {
	return mylib.SortStringByCharacter(s1) == mylib.SortStringByCharacter(s2)
}
