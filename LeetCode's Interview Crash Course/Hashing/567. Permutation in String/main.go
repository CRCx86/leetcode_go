package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkInclusion("abc", "bbbca"))
	fmt.Println(checkInclusion("ab", "ab"))
	fmt.Println(checkInclusion("adc", "dcda"))
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("abcdxabcde", "abcdeabcdx"))
}

func checkInclusion(s1 string, s2 string) bool {

	counts := make(map[uint8]int)
	for i := range s1 {
		counts[s1[i]]++
	}

	n := len(s1)
	for right := n - 1; right < len(s2); right++ {
		c := make(map[uint8]int)
		for j := right - n + 1; j <= right; j++ {
			c[s2[j]]++
		}
		for k := range counts {
			if _, ok := c[k]; ok {
				c[k] -= counts[k]
				if c[k] == 0 {
					delete(c, k)
				}
			}
		}
		if len(c) == 0 {
			return true
		}

	}

	return false
}
