package main

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}))
}

func asteroidCollision(asteroids []int) []int {

	stack := make([]int, 0)

	for i := 0; i < len(asteroids); i++ {
		if len(stack) == 0 {
			stack = append(stack, asteroids[i])
		} else {
			pop := stack[len(stack)-1]
			if pop > 0 && asteroids[i] > 0 || pop < 0 && asteroids[i] < 0 {
				stack = append(stack, asteroids[i])
			} else {
				last := asteroids[i]
				for len(stack) > 0 && pop > 0 && last < 0 {
					if mod(pop) > mod(last) {
						last = pop
					} else if mod(pop) == mod(last) {
						last = 0
					}
					stack = stack[:len(stack)-1]
					if len(stack) > 0 {
						pop = stack[len(stack)-1]
					}
				}
				if last != 0 {
					stack = append(stack, last)
				}
			}
		}
	}

	return stack
}

func mod(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
