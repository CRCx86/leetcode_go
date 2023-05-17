package main

import "fmt"

func main() {
	/**
	* Your MyQueue object will be instantiated and called as such:*/
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param3 := obj.Peek()
	param2 := obj.Pop()
	param4 := obj.Empty()
	fmt.Println(param2, param3, param4)
}

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{
		//stack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	//if len(this.stack) == 0 {
	//	return 0
	//}
	pop := this.stack[0]
	this.stack = this.stack[1:]
	return pop
}

func (this *MyQueue) Peek() int {
	//if len(this.stack) == 0 {
	//	return 0
	//}
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
