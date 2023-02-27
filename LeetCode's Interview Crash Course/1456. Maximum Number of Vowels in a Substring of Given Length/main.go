package main

import (
	"fmt"
	"math"
)

//	Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
//	Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

func main() {
	fmt.Println(maxVowels2("abciiidef", 3))
	fmt.Println(maxVowels("abciiidef", 9))
	fmt.Println(maxVowels("aeiou", 2))
	fmt.Println(maxVowels("leetcode", 3))
	fmt.Println(maxVowels("ramadan", 2))
}

func maxVowels(s string, k int) int {

	res := 0.0
	curr := 0.0
	left := 0

	for right := 0; right < k; right++ {
		if s[right] == 'a' || s[right] == 'e' || s[right] == 'i' || s[right] == 'o' || s[right] == 'u' {
			curr++
		}
	}

	res = curr
	for right := k; right < len(s); right++ {

		if s[right] == 'a' || s[right] == 'e' || s[right] == 'i' || s[right] == 'o' || s[right] == 'u' {
			curr++
		}

		if s[left] == 'a' || s[left] == 'e' || s[left] == 'i' || s[left] == 'o' || s[left] == 'u' {
			curr--
		}
		left++
		res = math.Max(res, curr)
	}

	return int(res)
}

func maxVowels2(s string, k int) int {

	curr := 0
	for _, c := range s[:k] {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			curr++
		}
	}

	max := curr
	for i, c := range s[k:] {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			curr++
		}
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			curr--
		}
		if curr > max {
			max = curr
		}
	}

	return max

}