package main

import (
	"fmt"
)

//	Given a string s, reverse the string according to the following rules:
//
//	All the characters that are not English letters remain in the same position.
//	All the English letters (lowercase or uppercase) should be reversed.
//
//	Return s after reversing it.

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
	fmt.Println(reverseOnlyLetters("7_28]"))
}

func reverseOnlyLetters(s string) string {

	bytes := []byte(s)

	left := 0
	right := len(bytes) - 1

	for left < right {

		if isAlpha(bytes[left]) && isAlpha(bytes[right]) {
			buff := bytes[left]
			bytes[left] = bytes[right]
			bytes[right] = buff
			left++
			right--
		} else if !isAlpha(bytes[left]) {
			left++
		} else {
			right--
		}

		//if bytes[left] < 65 || (bytes[left] > 90 && bytes[left] < 97) {
		//	left++
		//	continue
		//}
		//
		//if bytes[right] < 65 || (bytes[right] > 90 && bytes[right] < 97) {
		//	right--
		//	continue
		//}
		//
		//buff := bytes[left]
		//bytes[left] = bytes[right]
		//bytes[right] = buff
		//left++
		//right--
	}

	return string(bytes)
}

func isAlpha(any byte) bool {
	if (any >= 65 && any <= 90) || (any >= 97 && any <= 122) {
		return true
	}
	return false
}
