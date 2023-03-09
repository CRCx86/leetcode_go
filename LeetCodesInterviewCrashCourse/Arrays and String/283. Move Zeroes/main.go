package main

import "fmt"

//Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//Note that you must do this in-place without making a copy of the array.

func main() {
	ints := []int{0, 1, 0, 3, 12}
	moveZeroes(ints)
	fmt.Println(ints)
	ints = []int{0}
	moveZeroes(ints)
	fmt.Println(ints)
	ints = []int{1, 0, 1}
	moveZeroes(ints)
	fmt.Println(ints)
}

func moveZeroes(nums []int) {

	// slow
	//left := 0
	//right := 0
	//for left < len(nums) && right < len(nums) {
	//
	//	if nums[left] == 0 && nums[right] != 0 {
	//		nums[left], nums[right] = nums[right], nums[left]
	//		left++
	//	} else {
	//		if nums[left] != 0 && nums[right] == 0 {
	//			left = right
	//		}
	//		right++
	//	}
	//}

	// Fast - не моя идея
	ptr := 0
	for _, e := range nums {
		if e != 0 {
			nums[ptr] = e
			ptr++
		}
	}

	for i := ptr; i < len(nums); i++ {
		nums[i] = 0
	}

}
