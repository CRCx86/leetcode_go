package main

import "fmt"

// СПИСАНА!!! ДО ИДЕИ ДОГАДАЛСЯ, НО НЕ ДОКРУТИЛ
func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))

}

func lengthOfLongestSubstring(s string) int {

	left := 0
	ans := 0
	counts := make(map[uint8]int)
	for right := 0; right < len(s); right++ {
		counts[s[right]]++
		for counts[s[right]] > 1 {
			counts[s[left]]--
			left++
		}
		ans = _max(ans, right-left+1)
	}

	return ans
}

func _max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
