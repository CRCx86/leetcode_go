package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("b", "abc"))
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {

	m := len(s)
	if m == 0 {
		return true
	}

	j := 0
	for i := 0; i < len(t); i++ {
		if s[j] == t[i] {
			j++
		}
		if j == m {
			return true
		}
	}

	return false
}
