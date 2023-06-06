package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("2934"))
	fmt.Println(letterCombinations("22"))
}

var graph map[uint8][]string

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	numbers := make([]uint8, len(digits))
	for i := range digits {
		numbers[i] = digits[i] - 48
	}
	graph = make(map[uint8][]string)
	graph[2] = []string{"a", "b", "c"}
	graph[3] = []string{"d", "e", "f"}
	graph[4] = []string{"g", "h", "i"}
	graph[5] = []string{"j", "k", "l"}
	graph[6] = []string{"m", "n", "o"}
	graph[7] = []string{"p", "q", "r", "s"}
	graph[8] = []string{"t", "u", "v"}
	graph[9] = []string{"w", "x", "y", "z"}

	curr := make([]string, 0)
	ans := make([]string, 0)
	backtrack(curr, &ans, numbers, 0)

	return ans

}

func backtrack(curr []string, ans *[]string, numbers []uint8, i uint8) {

	if len(curr) == len(numbers) {
		*ans = append(*ans, strings.Join(curr, ""))
		return
	}

	for _, s := range graph[numbers[i]] {
		curr = append(curr, s)
		backtrack(curr, ans, numbers, i+1)
		curr = curr[0 : len(curr)-1]
	}
}
