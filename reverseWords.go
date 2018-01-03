/*
reverseWords: reverses the order of words of a sentence.
To run: go build reverseWords.go
		./reverseWords -Istr="your input string"
*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

var Istr = flag.String("Istr", "the sky is blue", "input string")

func main() {
	flag.Parse()
	fmt.Printf("input sentence: %s\n", *Istr)
	fmt.Printf("words in reversed order: %s\n", reverseWords(*Istr))
}

func reverseWords(s string) string {
	stack := []string{}
	words := strings.Split(s, " ") // split input into words delimited by space
	for _, w := range words {
		stack = append(stack, w) // push into stack
	}
	ans := []string{}
	for len(stack) > 0 {
		top := stack[len(stack)-1]   // top of stack
		stack = stack[:len(stack)-1] // pop
		ans = append(ans, top)
	}
	return strings.Join(ans, " ") // join into a string delimited by space
}
