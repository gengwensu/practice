/* Given a string, we can "shift" each of its letter to its successive letter, for example: "abc" -> "bcd". We can keep "shifting" which forms the sequence:

"abc" -> "bcd" -> ... -> "xyz"

Given a list of strings which contains only lowercase alphabets, group all strings that belong to the same shifting sequence.

For example, given: ["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"],
Return:

[
  ["abc","bcd","xyz"],
  ["az","ba"],
  ["acef"],
  ["a","z"]
]



Note: For the return value, each inner list's elements must follow the lexicographic order. */

package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}
	fmt.Printf("%v\n", groupStrings(input))
}

func groupStrings(strs []string) [][]string {
	ans := [][]string{}
	hm := map[string][]string{}
	for _, s := range strs {
		t := reduce(s)
		arr := hm[t]
		arr = append(arr, s)
		hm[t] = arr
	}

	for _, v := range hm {
		sort.Strings(v)
		ans = append(ans, v)
	}
	return ans
}

func reduce(s string) string {
	result := []rune{}
	rs := []rune(s)
	delta := rs[0] - 'a'
	for _, r := range s {
		result = append(result, r-delta)
	}
	return string(result)
}
