package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))

	obj = Constructor(1, []int{})
	fmt.Println(obj.Add(-3))
	fmt.Println(obj.Add(-2))
	fmt.Println(obj.Add(-4))
	fmt.Println(obj.Add(0))
	fmt.Println(obj.Add(4))

	obj = Constructor(2, []int{0})
	fmt.Println(obj.Add(-1))
	fmt.Println(obj.Add(1))
	fmt.Println(obj.Add(-2))
	fmt.Println(obj.Add(-4))
	fmt.Println(obj.Add(3))
}

/**
* Your KthLargest object will be instantiated and called as such:
* obj := Constructor(k, nums);
* param_1 := obj.Add(val);
 */

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {

	h := &IntHeap{}
	for _, n := range nums {
		*h = append(*h, n)
	}

	heap.Init(h)
	for i := len(nums); i > k; i-- {
		heap.Pop(h)
	}

	return KthLargest{
		h: h,
		k: k,
	}
}

func (t *KthLargest) Add(val int) int {

	heap.Push(t.h, val)
	if t.h.Len() > t.k {
		heap.Pop(t.h)
	}

	return (*t.h)[0]
}

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
