package main

import "fmt"

func main() {
	fmt.Println(longestSubstring("aaabb", 3))
}

func longestSubstring(s string, k int) int {

	res := 0
	left := 0
	m := make(map[uint8]uint, k+1)
	for right := 0; right < len(s); right++ {

		m[s[right]]++

		for len(m) > k {
			m[s[left]]--
			if m[s[left]] == 0 {
				delete(m, s[left])
			}
			left++
		}

		res = max(res, right-left+1)
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
