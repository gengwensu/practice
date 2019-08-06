/**
Given an array [a[i]], start from a[0] and take successive steps.
Each step will jump to i+a[i]. Write a function to calculate the number of steps it takes
to jump out of bound.
example a=[2,3,-1,2,-1] will take 5 steps to jump out of bound. return -1 if jumping out of
board is not possible
**/
package main

import "fmt"

func main() {
	//a := []int{2, 3, -1, 2, -1}
	//a := []int{1, -1, 1, -1, 1}
	//a := []int{2, 3, -1, 2, -1}
	a := []int{2, 3, -3, 2, -1}

	fmt.Printf("Step to jump out of bound - %d\n", arrayJump(a))
}

var visited []bool

func arrayJump(arr []int) int {
	visited = make([]bool, len(arr))
	return recursiveJump(arr, 0)
}

func recursiveJump(arr []int, position int) int {
	fmt.Printf("in recursiveJump, position - %d\n", position)
	if visited[position] { //trapped can't jump out of bound
		return -1
	}
	visited[position] = true

	nextp := position + arr[position]
	fmt.Printf("nextp - %d\n", nextp)
	if nextp < 0 || nextp >= len(arr) { //last step. base case
		return 1
	}

	step := recursiveJump(arr, nextp) //make the next jump
	if step == -1 {
		return -1
	}
	return step + 1
}
