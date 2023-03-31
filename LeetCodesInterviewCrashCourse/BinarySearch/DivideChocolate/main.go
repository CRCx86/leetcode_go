package main

import "fmt"

func main() {
	fmt.Println(maximizeSweetness([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(maximizeSweetness([]int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 8))
	fmt.Println(maximizeSweetness([]int{1, 2, 2, 1, 2, 2, 1, 2, 2}, 2))
}

func maximizeSweetness(sweetness []int, k int) int {

	left := 1
	right := 0
	for _, n := range sweetness {
		right += n
	}

	for left <= right {
		mid := left + (right-left)/2
		if check(mid, sweetness, k) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}

func check(mid int, sweetness []int, k int) bool {

	chunks := 0
	sum := 0

	for i := 0; i < len(sweetness); i++ {
		sum += sweetness[i]
		if sum >= mid {
			chunks++
			sum = 0
		}
	}

	return chunks <= k
}
