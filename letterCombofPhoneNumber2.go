/*
letterCombofPhoneNumber: a function that prints all combination of a phone number.
*/
package main

import (
	"flag"
	"fmt"
)

var Istr1 = flag.String("Istr1", "234", "input string")

func main() {
	flag.Parse()
	fmt.Printf("string1: %s\n", *Istr1)
	fmt.Printf("Letter combination of %s are: %v\n", *Istr1, letterCombinations(*Istr1))
}

func letterCombinations(digits string) []string {
	hm := map[rune]string{
		'0': "",
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ans := []string{}
	for _, d := range digits {
		out := []string{}
		for _, r := range hm[d] {
			if len(ans) == 0 {
				out = append(out, string(r))
			} else {
				for _, s := range ans {
					out = append(out, s+string(r))
				}
			}
		}
		ans = out
	}
	return ans
}
