package main

import "fmt"

// СПИСАНА!!! ДО ИДЕИ ДОГАДАЛСЯ, НО НЕ ДОКРУТИЛ
func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("aab"))

}

func lengthOfLongestSubstring(s string) int {

	max := 0
	left := 0
	counts := make(map[uint8]int)

	for right := 0; right < len(s); right++ {
		counts[s[right]]++

		for counts[s[right]] > 1 {
			counts[s[left]]--
			left++
		}

		ans := right - left + 1
		if ans > max {
			max = ans
		}

	}

	return max
}
