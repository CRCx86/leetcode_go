package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyStack struct {
	queue list.List
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue.PushBack(x)
}

func (this *MyStack) Pop() int {
	back := this.queue.Back()
	this.queue.Remove(back)
	return back.Value.(int)
}

func (this *MyStack) Top() int {
	return this.queue.Back().Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.queue.Len() == 0
}
