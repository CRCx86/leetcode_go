package main

import "fmt"

//	Given a 0-indexed string word and a character ch, reverse the segment of word that starts at index 0 and ends at the index of the first occurrence of ch (inclusive).
//	If the character ch does not exist in word, do nothing.
//
//	For example, if word = "abcdefd" and ch = "d", then you should reverse the segment that starts at 0 and ends at 3 (inclusive). The resulting string will be "dcbaefd".
//	Return the resulting string.

func main() {

	fmt.Println(reversePrefix("abcdefd", 'd'))
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))

}

func reversePrefix(word string, ch byte) string {

	bytes := make([]byte, len(word))
	start := 0
	for i, s := range word {
		bytes[i] = byte(s)
		if bytes[i] == ch && start == 0 {
			start = i + 1
		}
	}

	if start == 0 {
		return word
	}

	left := 0
	right := start - 1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}

	return string(bytes)
}
