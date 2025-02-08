package main

import "fmt"

func main() {
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
	// prefix: 7, 11, 14, 23, 24, 32, 37, 39, 45
	fmt.Println(getAverages([]int{100000}, 0))
	fmt.Println(getAverages([]int{8}, 100000))
	fmt.Println(getAverages([]int{40527, 53696, 10730, 66491, 62141, 83909, 78635, 18560}, 2))
}

func getAverages(nums []int, k int) []int {

	if k == 0 {
		return nums
	}

	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if i-k == 0 && i+k < n {
			ans[i] = prefix[i+k] / (2*k + 1)
			continue
		}
		if i-k > 0 && i+k < n {
			ans[i] = (prefix[i+k] - prefix[i-k-1]) / (2*k + 1)
			continue
		}
		ans[i] = -1
	}

	return ans
}
