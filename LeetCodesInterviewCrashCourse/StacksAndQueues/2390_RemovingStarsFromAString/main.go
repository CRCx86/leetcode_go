package main

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
}

func removeStars(s string) string {

	if len(s) == 0 {
		return ""
	}

	ans := make([]int32, 0)
	for _, c := range s {
		if c != '*' {
			ans = append(ans, c)
		} else {
			ans = ans[:len(ans)-1]
		}
	}

	return string(ans)
}
