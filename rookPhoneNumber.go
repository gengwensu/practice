/* Question: We are interested in getting 7 digit phone number by watching a chess piece (rook) traverse a board with numbers on each board position.

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
	"strconv"
)

var Istart = flag.Int("Istart", 4, "input string")
var hm map[int][]int

func main() {
	flag.Parse()
	fmt.Printf("start position: %d\n", *Istart)
	hm = map[int][]int{}
	hm[1] = []int{2, 3, 4, 7}
	hm[2] = []int{1, 3, 5, 8}
	hm[3] = []int{1, 2, 6, 9}
	hm[4] = []int{1, 5, 6, 7}
	hm[5] = []int{2, 4, 6, 8}
	hm[6] = []int{3, 4, 5, 9}
	hm[7] = []int{1, 4, 8, 9}
	hm[8] = []int{2, 5, 7, 9}
	hm[9] = []int{3, 6, 7, 8}
	digit := 3
	fmt.Printf("Letter combination of %s are: %v\n", generatePhoneNumber(*Istart, digit))
}

func generatePhoneNumber(start int, digit int) []string {
	if digit == 0 { //base case
		return []string{}
	}

	ans := []string{}
	for _, e := range hm[start] {
		for _, next := range generatePhoneNumber(e, digit-1) {
			ans = append(ans, strconv.Itoa(start)+strconv.Itoa(e)+next)
		}
	}
	return ans
}
