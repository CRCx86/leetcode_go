package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}

func maxProfit(prices []int) int {

	n := len(prices)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i := n - 1; i >= 0; i-- {
		for holding := 0; holding < 2; holding++ {
			ans := dp[i+1][holding]
			if holding == 0 {
				ref := -prices[i] + dp[i+1][1]
				if ref > ans {
					ans = ref
				}
			} else {
				ref := prices[i] + dp[i+2][0]
				if ref > ans {
					ans = ref
				}
			}

			dp[i][holding] = ans
		}
	}

	return dp[0][0]

}
