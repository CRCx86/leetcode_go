package main

import "fmt"

func main() {

	fmt.Println(divisorGame(2))
	fmt.Println(divisorGame(3))

}

func divisorGame(n int) bool {

	if n%2 == 0 {
		return true
	}
	return false

}
