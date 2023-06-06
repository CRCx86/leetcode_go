package main

import "fmt"

// Разбор задачи

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	//fmt.Println(longestCommonSubsequence("abc", "abc"))
	//fmt.Println(longestCommonSubsequence("abc", "def"))
}

var txt1 string
var txt2 string

func longestCommonSubsequence(text1 string, text2 string) int {

	return bottomUp(text1, text2)

	//txt1 = text1
	//txt2 = text2
	//memo := make([][]int, len(text1))
	//for i := range memo {
	//	memo[i] = make([]int, len(text2))
	//	for j := range memo[i] {
	//		memo[i][j] = -1
	//	}
	//}
	//
	//return topDown(0, 0, memo)
}

func topDown(i, j int, memo [][]int) int {

	if i == len(txt1) || j == len(txt2) {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	ans := 0
	if txt1[i] == txt2[j] {
		ans = 1 + topDown(i+1, j+1, memo)
	} else {
		left := topDown(i+1, j, memo)
		right := topDown(i, j+1, memo)
		if right > left {
			left = right
		}
		ans = left
	}

	memo[i][j] = ans
	return ans
}

func bottomUp(text1 string, text2 string) int {

	n := len(text1)
	m := len(text2)
	dp := make([][]int, n+1)
	for j := range dp {
		dp[j] = make([]int, m+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				if dp[i+1][j] > dp[i][j+1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j+1]
				}
			}
		}
	}

	return dp[0][0]
}
