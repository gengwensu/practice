/*
print all prime number up to the input
*/

package main

import (
	"flag"
	"fmt"
)

var INum = flag.Int("Num", 1000, "find prime less than input int")

func main() {
	flag.Parse()
	fmt.Printf("For input number %v, prime : 1", *INum)
	for i := 2; i < *INum; i++ {
		if IsPrime(i) {
			fmt.Printf(", %v", i)
		}
	}
	fmt.Printf("\n")
}

func IsPrime(n int) bool {
	for j := 2; j*j <= n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}
