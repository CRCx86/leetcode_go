package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}

func frequencySort(s string) string {

	counts := make(map[rune]int)
	for _, v := range s {
		counts[v]++
	}

	keys := make([]rune, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return counts[keys[i]] > counts[keys[j]]
	})

	ans := make([]rune, 0, len(s))
	for k := range keys {
		for i := 0; i < counts[keys[k]]; i++ {
			ans = append(ans, keys[k])
		}
	}

	return string(ans)
}
