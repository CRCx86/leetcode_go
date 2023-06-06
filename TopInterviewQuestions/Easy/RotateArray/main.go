package main

func main() {
	rotate([]int{1, 2}, 2)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 3)
	rotate([]int{-1}, 2)
	rotate([]int{1, 2}, 3)
}

func rotate(nums []int, k int) {

	k = k % len(nums)
	n := len(nums)
	if k > 0 {
		left := 0
		right := n - 1
		for left < right {
			val := nums[left]
			nums[left] = nums[right]
			nums[right] = val
			left++
			right--
		}

		left = 0
		right = k - 1
		for left < right {
			val := nums[left]
			nums[left] = nums[right]
			nums[right] = val
			left++
			right--
		}

		left = k
		right = n - 1
		for left < right {
			val := nums[left]
			nums[left] = nums[right]
			nums[right] = val
			left++
			right--
		}
	}
}
