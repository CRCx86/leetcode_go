package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {

	counts := make(map[uint8]int)
	for i := range ransomNote {
		counts[ransomNote[i]]++
	}

	for i := range magazine {
		if counts[magazine[i]] > 0 {
			counts[magazine[i]]--

			if counts[magazine[i]] == 0 {
				delete(counts, magazine[i])
			}

			if len(counts) == 0 {
				return true
			}
		}
	}

	if len(counts) == 0 {
		return true
	}

	return false
}
