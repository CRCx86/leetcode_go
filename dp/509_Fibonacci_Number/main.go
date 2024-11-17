package main

import "fmt"

func main() {

	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(10))

}

func fib(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	f := make([]int, 3)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i%3] = f[(i-2)%3] + f[(i-1)%3]
	}

	return f[(n-1)%3]
}
