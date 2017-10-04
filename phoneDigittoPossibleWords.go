/*
letterCombofPhoneNumber: a function that prints all legitimate words of a phone number against a dictionary.
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
	td := *Istr1
	ans := []string{}
	dict := []string{"dog", "dogs", "cat"}
	//find all substring of input string excluding '0' and '1'
	l := len(td)
	for i := 1; i <= l; i++ { // length i substring
		for j := 0; j <= l-i; j++ { //starting from j
			for _, r := range letterCombinations(td[j : i+j]) {
				ans = append(ans, r)
			}
		}
	}
	fmt.Printf("Legitimate words for phone digit %s are: \n", td)
	for _, w := range ans {
		for _, x := range dict {
			if w == x {
				fmt.Printf(" %s", w)
			}
		}
	}
	fmt.Println(" ")
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
