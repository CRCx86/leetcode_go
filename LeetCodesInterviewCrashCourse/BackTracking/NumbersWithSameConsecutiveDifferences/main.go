package main

import (
	"fmt"
)

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
}

// Эта задача наебалово и не решается с помощью backtracking
// она решается через dfs
func numsSameConsecDiff(n int, k int) []int {

	if n == 1 {
		return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	ans := make([]int, 0)
	for i := 1; i < 10; i++ {
		dfs(&ans, n-1, k, i)
	}
	return ans
}

func dfs(ans *[]int, n int, k int, curr int) {

	if n == 0 {
		*ans = append(*ans, curr)
		return
	}

	nexts := make([]int, 0)
	num := curr % 10
	nexts = append(nexts, num+k)
	if k != 0 {
		nexts = append(nexts, num-k)
	}
	for _, next := range nexts {
		if 0 <= next && next < 10 {
			newNext := curr*10 + next
			dfs(ans, n-1, k, newNext)
		}
	}
}
