package main

import "fmt"

func main() {

	fmt.Println(countElements([]int{1, 2, 3}))
	fmt.Println(countElements([]int{1, 1, 3, 3, 5, 5, 7, 7}))

}

func countElements(arr []int) int {

	buff := make(map[int]int)
	for _, e := range arr {
		buff[e]++
	}

	count := 0
	for _, e := range arr {
		if buff[e] != 0 && buff[e+1] != 0 {
			count++
		}
	}

	return count

}
