/***
Given an array arr[] of integers, find out the maximum difference between arr[j]-arr[i] for j>i and
arr[j]>arr[i].

Examples: If array is [2, 3, 10, 6, 4, 8, 1] then returned value should be 8
(Diff between 10 and 2). If array is [ 7, 9, 5, 6, 3, 2 ] then returned value should be 2
(Diff between 7 and 9)
***/
package main

import "fmt"

func main() {
	test1 := []int{2, 3, 10, 1, 4, 8, 11, 1}
	test2 := []int{7, 9, 5, 6, 3, 2}
	test3 := []int{9, 7, 5, 3}
	fmt.Printf("Max difference for test1 is %d\n", maxDifferences(test1))
	fmt.Printf("Max difference for test2 is %d\n", maxDifferences(test2))
	fmt.Printf("Max difference for test3 is %d\n", maxDifferences(test3))
}

/***
Find the difference between the adjacent elements of the array.
Now this problems turns into finding the maximum sum subarray of this difference array.
***/
func maxDifferences(arr []int) int {
	currMin := arr[0]
	maxDiff := arr[1] - arr[0]

	for i := 1; i < len(arr); i++ {
		if diff := arr[i] - currMin; diff > maxDiff {
			maxDiff = diff
		}
		if arr[i] < currMin {
			currMin = arr[i]
		}
	}
	return maxDiff
}
