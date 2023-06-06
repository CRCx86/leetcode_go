package main

import "fmt"

func main() {
	fmt.Println(isPoly("racecar"))
	fmt.Println(isPoly("aceba"))
}

func isPoly(s string) bool {

	l := len(s)

	if l <= 1 {
		return true
	}

	left := 0
	right := l - 1

	for left < right {

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
