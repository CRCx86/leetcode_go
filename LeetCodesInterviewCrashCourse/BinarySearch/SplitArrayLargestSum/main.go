package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println(splitArray([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(splitArray([]int{1, 4, 4}, 3))
	fmt.Println(splitArray([]int{2, 3, 1, 2, 4, 3}, 5))
}

// Эту задачу разобрал. До конца не понял как догадаться, что:
// 1. откуда взялся элемент max
// 2. почему нужно присваивание min = mid
func splitArray(nums []int, k int) int {

	right := 0
	max := math.MinInt
	for _, n := range nums {
		right += n
		if n > max {
			max = n
		}
	}

	min := 0
	left := max
	for left <= right {
		mid := left + (right-left)/2
		if check(mid, nums, k) {
			right = mid - 1
			min = mid
		} else {
			left = mid + 1
		}
	}

	return min
}

// ключевая идея: получилось разбить массив на k-раз со значением mid, пробуем чутка увеличить значение mid,
// т.е двигаем left
// не получилось - двигаем right
func check(mid int, nums []int, k int) bool {

	chunks := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > mid {
			chunks++
			sum = nums[i]
		}
	}

	if sum > mid {
		chunks++
	}

	return chunks < k
}
