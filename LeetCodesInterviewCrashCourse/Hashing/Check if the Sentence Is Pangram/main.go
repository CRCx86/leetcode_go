package main

import "fmt"

func main() {

	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))

}

func checkIfPangram(sentence string) bool {

	//m := make(map[int32]int)
	//for _, c := range sentence {
	//	m[c] += 1
	//}
	//
	//return 26 == len(m)

	seen := 0
	for _, c := range sentence {

		left := c - 'a'

		find := 1 << left

		seen |= find

	}

	total := 1 << 26

	return seen == total-1

}
