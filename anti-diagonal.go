/***
Give a N*N square matrix, return an array of its anti-diagonals. Look at the example for more details.
Example:
Input:

1 2 3
4 5 6
7 8 9

Return the following :

[
  [1],
  [2, 4],
  [3, 5, 7],
  [6, 8],
  [9]
]
***/

package main

import "fmt"

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	r := [][]int{}
	row := 0
	for i := 0; i < len(a[0]); i++ {
		t := []int{}
		col := i
		for col >= 0 {
			t = append(t, a[row][col])
			row++
			col--
		}
		r = append(r, t)
		row = 0
	}
	col := len(a[0]) - 1
	for j := 1; j < len(a[0]); j++ {
		t := []int{}
		row = j
		for row < len(a[0]) {
			t = append(t, a[row][col])
			row++
			col--
		}
		r = append(r, t)
		col = len(a[0]) - 1
	}
	fmt.Printf("The result is %v\n", r)
}
