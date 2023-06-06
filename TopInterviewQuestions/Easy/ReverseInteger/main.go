package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(reverse(123))
	//fmt.Println(reverse(-123))
	//fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {

	res := 0
	if x > 0 {
		for x > 0 {
			div := x % 10
			if res > 0 {
				res *= 10
				res += div
			} else {
				res = div
			}
			x = x / 10
		}

	} else {
		x = -x
		for x > 0 {
			div := x % 10
			if res > 0 {
				res *= 10
				res += div
			} else {
				res = div
			}
			x = x / 10
		}
		res = -res
	}

	if res >= math.MaxInt32 || res <= math.MinInt32 {
		return 0
	}

	return res
}
