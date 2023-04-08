package main

import (
	"fmt"
)

// Разбор задачи

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

var pprices []int
var memo [][][]int

func maxProfit(k int, prices []int) int {

	return bottomDown(k, prices)

	//pprices = prices
	//memo = make([][][]int, len(prices))
	//for i := range memo {
	//	memo[i] = make([][]int, 2)
	//	for j := range memo[i] {
	//		memo[i][j] = make([]int, k+1)
	//		for l := range memo[i][j] {
	//			memo[i][j][l] = -1
	//		}
	//	}
	//}
	//
	//return topDown(0, 0, k)
}

func topDown(i int, holding int, remain int) int {
	if i == len(pprices) || remain == 0 {
		return 0
	}

	if memo[i][holding][remain] != -1 {
		return memo[i][holding][remain]
	}

	ans := topDown(i+1, holding, remain)
	if holding == 1 {
		ref := pprices[i] + topDown(i+1, 0, remain-1)
		if ref > ans {
			ans = ref
		}
	} else {
		ref := -pprices[i] + topDown(i+1, 1, remain)
		if ref > ans {
			ans = ref
		}
	}

	memo[i][holding][remain] = ans
	return ans
}

func bottomDown(k int, prices []int) int {

	n := len(prices)
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}

	for i := n - 1; i >= 0; i-- {
		for remain := 1; remain <= k; remain++ {
			for holding := 0; holding < 2; holding++ {
				ans := dp[i+1][holding][remain]
				if holding == 1 {
					ref := prices[i] + dp[i+1][0][remain-1]
					if ref > ans {
						ans = ref
					}
				} else {
					ref := -prices[i] + dp[i+1][1][remain]
					if ref > ans {
						ans = ref
					}
				}

				dp[i][holding][remain] = ans
			}
		}
	}

	return dp[0][0][k]
}
