/*
letterCombofPhoneNumber: a function that prints all legitimate words of a phone number against a dictionary.
*/
package main

import (
	"flag"
	"fmt"
)

var Istr1 = flag.String("Istr1", "234", "input string")
var pMap = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"} //phone digit mapping

func main() {
	flag.Parse()
	fmt.Printf("string1: %s\n", *Istr1)

	iStr := *Istr1
	ans := []string{}
	dict := []string{"dog", "dogs", "cat"}
	//find all substring of input string excluding '0' and '1'

	for sl := 2; sl <= len(iStr); sl++ { // substring length
		for i := 0; i <= len(iStr)-sl; i++ { //starting from j
			for _, r := range letterCombinations(iStr[i:i+sl], "") {
				ans = append(ans, r)
			}
		}
	}
	fmt.Printf("Words for phone digit %s are: \n", iStr)
	for _, w := range ans {
		for _, x := range dict {
			if w == x {
				fmt.Printf(" %s", w)
			}
		}
	}
	fmt.Println(" ")
}

func letterCombinations(digits, prefix string) []string {
	ret := []string{}
	if len(digits) == 1 { //last digit
		for _, t := range pMap[digits[0]-'0'] {
			ret = append(ret, prefix+string(t))
		}
	} else { // not the lat digit
		for _, t := range pMap[digits[0]-'0'] {
			ret = appendArray(ret, letterCombinations(digits[1:], prefix+string(t)))
		}
	}
	return ret
}

func appendArray(a1, a2 []string) []string {
	for _, s := range a2 {
		a1 = append(a1, s)
	}
	return a1
}
