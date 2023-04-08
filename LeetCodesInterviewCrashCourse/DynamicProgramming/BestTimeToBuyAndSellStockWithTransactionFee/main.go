package main

import "fmt"

func main() {
	//fmt.Println(maxProfit([]int{1, 3, 2}, 2))
	fmt.Println(maxProfit([]int{1, 3, 2, 8}, 2))

	//fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	//fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

var memo [][]int
var arrPrices []int
var intFee int

func maxProfit(prices []int, fee int) int {

	return dpBU(prices, fee)

	//arrPrices = prices
	//intFee = fee
	//
	//memo = make([][]int, len(prices))
	//for i := range memo {
	//	memo[i] = make([]int, 2)
	//	for j := range memo[i] {
	//		memo[i][j] = -1
	//	}
	//}
	//
	//return dpTD(0, 0)
}

func dpTD(i int, holding int) int {

	if len(arrPrices) == i {
		return 0
	}

	if memo[i][holding] != -1 {
		return memo[i][holding]
	}

	ans := dpTD(i+1, holding)
	if holding == 1 {
		ref := arrPrices[i] + dpTD(i+1, 0)
		if ref > ans {
			ans = ref
		}
	} else {
		ref := -arrPrices[i] + dpTD(i+1, 1) - intFee
		if ref > ans {
			ans = ref
		}
	}

	memo[i][holding] = ans
	return ans
}

func dpBU(prices []int, fee int) int {

	n := len(prices)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	for i := n - 1; i >= 0; i-- {
		for holding := 0; holding < 2; holding++ {
			ans := dp[i+1][holding]
			if holding == 1 {
				ref := prices[i] + dp[i+1][0]
				if ref > ans {
					ans = ref
				}
			} else {
				ref := -prices[i] + dp[i+1][1] - fee
				if ref > ans {
					ans = ref
				}
			}
			dp[i][holding] = ans
		}
	}

	return dp[0][0]
}
