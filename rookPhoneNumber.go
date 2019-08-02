/* Question: We are stringerested in getting 7 digit phone number by watching a chess piece (rook) traverse a board with numbers on each board position.

Given a 3 by 3 board that looks like this.

1 2 3
4 5 6
7 8 9

And assuming that a rook starts on position 4, write a function which provides all of the combinations of 7 digit numbers which start with 4.

For example:

456-3214 would be generated when the rook moves
{ right one, right one, up one, left one, left one, down one }

464-6464 would be generated when the rook moves
{right two, left two, right two, left two, right two, left two } */

package main

import (
	"flag"
	"fmt"
)

var Istart = flag.String("Istart", "4", "input string")
var Idigit = flag.Int("Idigit", 7, "Number of digits to generate")
var hm map[string][]string

func main() {
	flag.Parse()
	fmt.Printf("start position: %s\n", *Istart)
	fmt.Printf("Digit is %d\n", *Idigit)
	hm = map[string][]string{}
	hm["1"] = []string{"2", "3", "4", "7"}
	hm["2"] = []string{"1", "3", "5", "8"}
	hm["3"] = []string{"1", "2", "6", "9"}
	hm["4"] = []string{"1", "5", "6", "7"}
	hm["5"] = []string{"2", "4", "6", "8"}
	hm["6"] = []string{"3", "4", "5", "9"}
	hm["7"] = []string{"1", "4", "8", "9"}
	hm["8"] = []string{"2", "5", "7", "9"}
	hm["9"] = []string{"3", "6", "7", "8"}

	fmt.Printf("start - %s, generated: %v\n", *Istart, generatePhoneNumber(*Istart, *Idigit))
}

func generatePhoneNumber(start string, digit int) []string {
	if digit == 1 { //base case
		return []string{start}
	}

	ans := []string{}
	for _, e := range hm[start] {
		for _, next := range generatePhoneNumber(e, digit-1) {
			ans = append(ans, start+next)
		}
	}
	return ans
}
