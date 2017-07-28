/*
Comma: Non-recursive version of comma, using bytes,Buffer instead of string
concatenation
*/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

var Istr = flag.String("Istr", "abc", "input str for IsUnique")

func main() {
	flag.Parse()
	fmt.Printf("Input string %s\n", *Istr)
	fmt.Println(myComma(*Istr))
}

func myComma(s string) string {
	switch strings.Count(s, ".") {
	case 1:
		sl := strings.Split(s, ".") //split the string into two substring
		return intComma(sl[0]) + "." + intComma(sl[1])
	case 0:
		return intComma(s)
	default:
		return "illegal format"
	}
}

func intComma(s string) string {
	n := len(s) //length, assuming all numbers
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	rem := n % 3
	if rem > 0 {
		buf.WriteString(s[:rem])
	}
	for i := rem; i < n; i += 3 {
		if i > 0 {
			buf.WriteByte(',')
		}

		buf.WriteString(s[i : i+3])

	}
	return buf.String()
}
