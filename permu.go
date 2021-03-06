/*
This code counts all permutations of a string
*/

package main

import (
	"flag"
	"fmt"
)

var Istr = flag.String("Istr", "abc", "input for str permu")

func main() {
	flag.Parse()
	fmt.Printf("Input string %s\n", *Istr)
	permutation(*Istr, nil)
}

func permutation(str string, prefix []byte) {
	if len(str) == 0 { //no more char to pick
		fmt.Println(string(prefix))
	} else {
		for i := 0; i < len(str); i++ {
			p := append(prefix, str[i]) //move to prefix
			rem := str[:i] + str[i+1:]
			//fmt.Printf("rem is %s and prefix is %s\n", rem, string(p))
			permutation(rem, p)
		}
	}
}
