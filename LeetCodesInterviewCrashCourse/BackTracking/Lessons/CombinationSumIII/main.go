package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}

func combinationSum3(k int, n int) [][]int {

	ans := make([][]int, 0)
	path := make([]int, 0)
	backtrack(&ans, path, 1, 0, n, k)
	return ans
}

func backtrack(ans *[][]int, path []int, start int, curr int, target int, k int) {

	if curr == target && len(path) == k {
		dist := make([]int, len(path))
		copy(dist, path)
		*ans = append(*ans, dist)
		return
	}

	for i := start; i < 10; i++ {
		if curr+i <= target {
			path = append(path, i)
			backtrack(ans, path, i+1, curr+i, target, k)
			path = path[0 : len(path)-1]
		}
	}

}
