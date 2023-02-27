package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber(nums []int) int {

	numbers := make(map[int]bool)
	for _, n := range nums {
		numbers[n] = true
	}

	for i := 0; i <= len(nums); i++ {
		if !numbers[i] {
			return i
		}
	}

	return 0

}
