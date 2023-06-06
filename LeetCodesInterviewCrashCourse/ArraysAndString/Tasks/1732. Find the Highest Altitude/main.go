package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}

func largestAltitude(gain []int) int {

	max := 0
	//prefixes := make([]int, len(gain)+1)
	//for i := 1; i <= len(gain); i++ {
	//	prefixes[i] = prefixes[i-1] + gain[i-1]
	//}
	//
	//for i := 0; i < len(prefixes); i++ {
	//	max = math.Max(max, float64(prefixes[i]))
	//}

	sum := 0
	for _, e := range gain {
		sum += e
		if sum > max {
			max = sum
		}
	}

	return max
}
