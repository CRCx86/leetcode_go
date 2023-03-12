package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
	fmt.Println(makeGood("s"))
}

func makeGood(s string) string {

	stack := make([]int32, 0)
	for _, str := range s {
		if len(stack) > 0 && math.Abs(float64(str-stack[len(stack)-1])) == 32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, str)
		}
	}

	return string(stack)
}
