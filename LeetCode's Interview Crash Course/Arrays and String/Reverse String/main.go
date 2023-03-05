package main

import "fmt"

// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte) {

	left := 0
	right := len(s) - 1

	for left < right {

		a := s[left]
		s[left] = s[right]
		s[right] = a
		left++
		right--
	}

}
