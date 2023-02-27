package main

import (
	"fmt"
)

func main() {
	var s string
	s = "a"
	fmt.Println(longestPoly(s))
	s = "ba"
	fmt.Println(longestPoly(s))
	s = "bb"
	fmt.Println(longestPoly(s))
	s = "babad"
	fmt.Println(longestPoly(s))
	s = "cbbd"
	fmt.Println(longestPoly(s))
	s = "babab"
	fmt.Println(longestPoly(s))
}

func longestPoly(s string) string {

	n := len(s)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][i] = 1
	}

	p := 1
	I := 0
	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			if s[i] == s[j] && dp[i+1][j-1] == j-i-1 {
				dp[i][j] = dp[i+1][j-1] + 2
				if dp[i][j] > p {
					p = dp[i][j]
					I = i
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	return s[I : I+p]

	//var k int
	//a := make([]string, p)
	//m := p - 1
	//for I < n && J >= 0 && dp[I][J] != 0 {
	//	if I == J {
	//		a[k] = string(s[I])
	//	} else {
	//		a[k] = string(s[I])
	//		a[m] = string(s[J])
	//	}
	//	I += 1
	//	J -= 1
	//	k++
	//	m--
	//}
	//
	//return strings.Join(a, "")
}
