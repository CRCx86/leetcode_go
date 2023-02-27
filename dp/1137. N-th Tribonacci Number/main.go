package main

import "fmt"

func main() {
	fmt.Println(tribonacci(0))
	fmt.Println(tribonacci(1))
	fmt.Println(tribonacci(2))
	fmt.Println(tribonacci(3))
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(5))
	fmt.Println(tribonacci(6))
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	f := make([]int, 4)
	f[0] = 0
	f[1] = 1
	f[2] = 1
	f[3] = 2
	for i := 4; i <= n; i++ {
		f[i%4] = f[(i-3)%4] + f[(i-2)%4] + f[(i-1)%4]
	}

	return f[n%4]

}
