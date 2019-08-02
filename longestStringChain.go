/* Given a list of words, each word consists of English lowercase letters.

Let's say word1 is a predecessor of word2 if and only if we can add exactly one letter anywhere in word1 to make it equal to word2.  For example, "abc" is a predecessor of "abac".

A word chain is a sequence of words [word_1, word_2, ..., word_k] with k >= 1, where word_1 is a predecessor of word_2, word_2 is a predecessor of word_3, and so on.

Return the longest possible length of a word chain with words chosen from the given list of words.



Example 1:

Input: ["a","b","ba","bca","bda","bdca"]
Output: 4
Explanation: one of the longest word chain is "a","ba","bda","bdca".



Note:

    1 <= words.length <= 1000
    1 <= words[i].length <= 16
    words[i] only consists of English lowercase letters. */
package main

import (
	"fmt"
	"strings"
)

var hm, tm map[string]int

func main() {
	//input := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	//input := []string{"ksqvsyq", "ks", "kss", "czvh", "zczpzvdhx", "zczpzvh", "zczpzvhx", "zcpzvh", "zczvh", "gr", "grukmj", "ksqvsq", "gruj", "kssq", "ksqsq", "grukkmj", "grukj", "zczpzfvdhx", "gru"}
	input := []string{"cefkp", "efkp", "pkacefkep", "pkacefkp", "pacefkp", "acefkp"}
	fmt.Printf("The longest string chain length is %d", longestStrChain(input))
}

func longestStrChain(words []string) int {
	if len(words) == 0 {
		return 0
	}

	hm = map[string]int{} // key is word, value is length of the string chain
	tm = map[string]int{}
	for _, w := range words {
		if _, ok := tm[w]; !ok {
			tm[w] = 1
		}
	}

	maxLen := 0

	for _, w := range words {
		if _, ok := hm[w]; !ok {
			if hm[w] = calStringChainLen(w, words); hm[w] > maxLen {
				maxLen = hm[w]
			}
			fmt.Printf("back from calSCL. hm[%s] is %d, maxLen is %d\n", w, hm[w], maxLen)
		}
	}
	return maxLen
}

func calStringChainLen(w string, words []string) int {
	fmt.Printf("in calSCL, w - %s, words - %v\n", w, words)
	if scl, ok := hm[w]; ok { // already calculated
		return scl
	}

	if len(w) < 2 {
		return 1
	}

	maxscl := 1
	sl := strings.Split(w, "")
	for i := 0; i < len(sl); i++ {
		tl := make([]string, len(sl)-1)
		copy(tl, sl[:i])
		copy(tl[i:], sl[i+1:])
		subword := strings.Join(tl, "")
		fmt.Printf("i - %d, subword %s\n", i, subword)
		scl := 0
		if _, ok := tm[subword]; ok {
			if scl, ok = hm[w]; !ok {
				scl = calStringChainLen(subword, words)
				hm[subword] = scl
			}
			scl++
			fmt.Printf("after ++, scl - %d, maxscl - %d\n", scl, maxscl)
			if scl > maxscl {
				maxscl = scl
			}
		}
	}
	hm[w] = maxscl
	return maxscl
}
