package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
	fmt.Println(minSetSize([]int{7, 7, 7, 7, 7, 7}))
}

type pair struct {
	num  int
	freq int
}

func minSetSize(arr []int) int {

	counts := make(map[int]int)
	for _, e := range arr {
		counts[e]++
	}

	pairs := make([]pair, 0)
	for k, v := range counts {
		pairs = append(pairs, pair{
			num:  k,
			freq: v,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	k := 0
	ans := 0
	for k < len(arr)/2 && ans < len(pairs) {
		k += pairs[ans].freq
		ans++
	}

	return ans
}
