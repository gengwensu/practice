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

var ret []string

func help(btn []string, str, digits string, idx int) {
	if idx == len(digits) {
		ret = append(ret, str)
		return
	}
	tmp := btn[digits[idx]-'0']
	for _, t := range tmp {
		help(btn, str+string(t), digits, idx+1)
	}
}

func letterCombinations(digits string) []string {
	btn := []string{" ", " ", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ret = []string{}
	if len(digits) > 0 {
		help(btn, "", digits, 0)
	}
	return ret
}
