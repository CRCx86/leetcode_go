package main

import "fmt"

// 322. Coin Change

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}

func coinChange(coins []int, amount int) int {

	max := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0

	// TODO: доразобрать задачу. Непонятно нихрена.
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				min := dp[i-coins[j]] + 1
				if dp[i] < min {
					min = dp[i]
				}
				dp[i] = min
			}
		}
	}

	return -1
}
