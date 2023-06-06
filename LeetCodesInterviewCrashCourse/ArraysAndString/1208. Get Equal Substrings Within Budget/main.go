package main

import (
	"fmt"
	"math"
)

//	You are given two strings s and t of the same length and an integer maxCost.
//
//	You want to change s to t. Changing the ith character of s to ith character of t costs |s[i] - t[i]|(i.e.,
//	the absolute difference between the ASCII values of the characters).
//
//	Return the maximum length of a substring of s that can be changed to be the same as the corresponding substring
//	of t with a cost less than or equal to maxCost. If there is no substring from s that can be changed to its corresponding substring from t, return 0.

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
	fmt.Println(equalSubstring("abcd", "cdef", 3))
	fmt.Println(equalSubstring("abcd", "acde", 0))
	fmt.Println(equalSubstring("krrgw", "zjxss", 19))
}

func equalSubstring(s string, t string, maxCost int) int {

	//maxCost -= int(math.Abs(float64(c - rune(t[i]))))
	res := 0.0
	left := 0
	for right := 0; right < len(s); right++ {
		maxCost -= int(math.Abs(float64(int(s[right]) - int(t[right]))))
		for maxCost < 0 {
			maxCost += int(math.Abs(float64(int(s[left]) - int(t[left]))))
			left++
		}
		res = math.Max(res, float64(right-left+1))
	}

	return int(res)
}
