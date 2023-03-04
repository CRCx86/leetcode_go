package main

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
	fmt.Println(maxNumberOfBalloons("leetcode"))
	fmt.Println(maxNumberOfBalloons("lloo"))
	fmt.Println(maxNumberOfBalloons("balon"))
}

func maxNumberOfBalloons(text string) int {

	ans := len(text)

	counts := make(map[uint8]int)
	balloon := make(map[uint8]int)
	balloon['b'] = 1
	balloon['a'] = 1
	balloon['l'] = 2
	balloon['o'] = 2
	balloon['n'] = 1

	for letter := range text {
		if balloon[text[letter]] > 0 {
			counts[text[letter]]++
		}
	}

	if len(balloon) != len(counts) {
		return 0
	}

	for k, v := range balloon {
		if buff := counts[k] / v; buff > 0 && counts[k] >= v {
			if buff < ans {
				ans = buff
			}
		} else {
			return 0
		}
	}

	if ans == len(text) {
		ans = 0
	}

	return ans
}
