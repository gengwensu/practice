/***
regexpMatch implements regular expression matching with support for '.' and '*'.
'.' Matches any single character.
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true
***/
package main

import (
	"flag"
	"fmt"
)

func main() {
	var istr = flag.String("istr", "aa", "input string")
	var mstr = flag.String("mstr", "a", "matching string")
	flag.Parse()
	s, p := *istr, *mstr
	fmt.Printf("input: %s, matching string: %s, isMatch returns ", s, p)
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	//fmt.Printf("input %s %s\n", s, p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	} // dp[i][j] is the state s[0:i] matches p[0:j]

	dp[0][0] = true
	for i := range p { // taking care of s="aab", p="c*a*b"
		if p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}

	for i := 0; i < m; i++ {
		//fmt.Printf("i = %d\n", i)
		for j := 0; j < n; j++ {
			//fmt.Printf("dp[%d][%d] is %t\n", i, j, dp[i][j])
			if p[j] != '*' {

				dp[i+1][j+1] = dp[i][j] && (p[j] == '.' || s[i] == p[j])
				//fmt.Printf("dp[%d][%d] is %t\n", i+1, j+1, dp[i+1][j+1])
			} else { // '*' case
				if p[j-1] != s[i] && p[j-1] != '.' { // repeats 0 times
					dp[i+1][j+1] = dp[i+1][j-1]
				} else { // repeats at least 1 time
					dp[i+1][j+1] = (dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1])
				}
				//fmt.Printf("dp[%d][%d] is %t\n", i+1, j+1, dp[i+1][j+1])
			}

		}
	}

	//fmt.Printf("dp[%d][%d] is %t\n",m,n,dp[m][n])
	return dp[m][n]
}
