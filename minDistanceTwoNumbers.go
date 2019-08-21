/*
Find the minimum distance between two numbers

Given an unsorted array arr[] and two numbers x and y, find the minimum distance between x and y in arr[]. The array might also contain duplicates. You may assume that both x and y are different and present in arr[].

Examples:
Input: arr[] = {1, 2}, x = 1, y = 2
Output: Minimum distance between 1 and 2 is 1.

Input: arr[] = {3, 4, 5}, x = 3, y = 5
Output: Minimum distance between 3 and 5 is 2.

Input: arr[] = {3, 5, 4, 2, 6, 5, 6, 6, 5, 4, 8, 3}, x = 3, y = 6
Output: Minimum distance between 3 and 6 is 4.

Input: arr[] = {2, 5, 3, 5, 4, 4, 2, 3}, x = 3, y = 2
Output: Minimum distance between 3 and 2 is 1. */
package main

import "fmt"

func main() {
	// input := []int{2, 5, 3, 5, 4, 4, 2, 3}
	// x, y := 3, 2
	input := []int{3, 5, 4, 2, 6, 5, 6, 6, 5, 4, 8, 3}
	x, y := 3, 6

	fmt.Printf("Min distance - %d\n", minDist(input, x, y))
}

func minDist(arr []int, x, y int) int {
	ix, iy := 0, 0
	min := len(arr)
	for i, num := range arr {
		if num == x {
			ix = i
		}
		if num == y {
			iy = i
		}
		if d := ix - iy; d > 0 {
			if d < min {
				min = d
			}
		} else if -d > 0 {
			if -d < min {
				min = -d
			}
		}
	}
	return min
}
