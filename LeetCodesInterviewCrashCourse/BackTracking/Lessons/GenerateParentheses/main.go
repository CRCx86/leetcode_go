package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {

	ans := make([]string, 0)
	curr := make([]string, 0)
	backtrack(&ans, n, curr, 0, 0)

	return ans
}

func backtrack(ans *[]string, n int, curr []string, start int, opened int) {

}
