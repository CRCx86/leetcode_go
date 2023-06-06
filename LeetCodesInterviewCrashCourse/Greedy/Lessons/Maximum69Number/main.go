package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))

	fmt.Println(maximum69Number2(9669))
	fmt.Println(maximum69Number2(9996))
	fmt.Println(maximum69Number2(9999))
}

func maximum69Number(num int) int {

	max := num

	nums := make([]int, 0)
	for num != 0 {
		i := num % 10
		nums = append(nums, i)
		num = num / 10
	}

	for i := len(nums) - 1; i >= 0; i-- {
		buff := nums[i]
		if nums[i] == 9 {
			nums[i] -= 3
		} else {
			nums[i] += 3
		}
		ans := 0
		for j := 0; j < len(nums); j++ {
			ans += nums[j] * int(math.Pow(10, float64(j)))
		}
		if ans > max {
			max = ans
		}
		nums[i] = buff
	}

	return max
}

func maximum69Number2(num int) int {

	index6 := -1
	buff := num
	index := 0
	for num != 0 {
		i := num % 10
		if i == 6 {
			index6 = index
		}
		num = num / 10
		index++
	}

	if index6 == -1 {
		return buff
	} else {
		return buff + 3*int(math.Pow(10, float64(index6)))
	}

}
