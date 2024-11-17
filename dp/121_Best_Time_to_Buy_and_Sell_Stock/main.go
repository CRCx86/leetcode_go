package main

import (
	"fmt"
	"math"
)

func main() {
	//maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(maxProfit([]int{7, 8}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {

	var res int
	n := len(prices)

	minRes := make([]int, n)
	minRes[0] = prices[0]
	for i := 1; i < n; i++ {
		minRes[i] = int(math.Min(float64(minRes[i-1]), float64(prices[i])))
		diff := prices[i] - minRes[i]
		if diff > res {
			res = diff
		}
	}

	return res
}
