package main

import "fmt"

//	Given a string s, reverse the order of characters in each word within
//	a sentence while still preserving whitespace and initial word order.

func main() {

	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("God Ding"))

}

func reverseWords(s string) string {

	i2 := len(s)
	left := 0
	a := make([]byte, i2)
	for i := 0; i < i2; i++ {
		a[i] = s[i]
		if s[i] == ' ' {
			right := i - 1
			for left < right {
				buff := a[left]
				a[left] = a[right]
				a[right] = buff
				left++
				right--
			}
			left = i + 1
		}
	}

	right := i2 - 1
	for left < right {
		buff := a[left]
		a[left] = a[right]
		a[right] = buff
		left++
		right--
	}

	return string(a)
}
