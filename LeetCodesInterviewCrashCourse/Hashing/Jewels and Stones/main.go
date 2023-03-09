package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(jewels string, stones string) int {

	counts := make(map[uint8]int)
	for i := range jewels {
		counts[jewels[i]]++
	}

	count := 0
	for i := range stones {
		if counts[stones[i]] > 0 {
			count++
		}
	}

	return count
}
