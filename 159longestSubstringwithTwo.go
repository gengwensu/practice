/* [LeetCode] 159. Longest Substring with At Most Two Distinct Characters



Given a string s , find the length of the longest substring t  that contains at most 2 distinct characters.

Example 1:

Input: "eceba"
Output: 3
Explanation: tis "ece" which its length is 3.

Example 2:

Input: "ccaabbb"
Output: 5
Explanation: tis "aabbb" which its length is 5. */
package main

import "fmt"

func main() {
	input := "ccaabbb"
	//input := "eceba"

	fmt.Printf("max length - %d\n", lengthOfLongestSubstringTwoDistinct(input))
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	arr := []rune(s)
	left := 0
	hm := map[rune]int{}
	max := 0
	for i, r := range arr {
		hm[r] = i
		for len(hm) > 2 {
			if hm[arr[left]] == left {
				delete(hm, arr[left])
			}
			left++
		}
		if d := i - left + 1; d > max {
			max = d
		}
	}
	return max
}
