package main

import "fmt"

func main() {
	fmt.Println(robotWithString("zza"))
	fmt.Println(robotWithString("bac"))
	fmt.Println(robotWithString("bdda"))
	fmt.Println(robotWithString("bydizfve"))
}

func robotWithString(s string) string {

	if len(s) == 0 {
		return s
	}
	//ans := make([]uint8, 0)
	stack := make([]uint8, 0)
	stack = append(stack, s[0])
	for i := 0; i < len(s); i++ {

		pop := stack[len(stack)-1]
		if s[i] > pop {
			stack = stack[:len(stack)-1]
			stack = append(stack, s[i])
			stack = append(stack, pop)
		} else {
			stack = append(stack, s[i])
		}

		//if s[i] != 'a' && len(ans) == 0 {
		//	stack = append(stack, s[i])
		//} else {
		//	ans = append(ans, s[i])
		//	for len(stack) > 0 {
		//		ans = append(ans, stack[len(stack)-1])
		//		stack = stack[:len(stack)-1]
		//	}
		//}

	}

	return string(stack)

}
