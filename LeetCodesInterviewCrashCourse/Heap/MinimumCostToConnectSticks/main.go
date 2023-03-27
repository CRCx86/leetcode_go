package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(connectSticks([]int{2, 4, 3}))
	fmt.Println(connectSticks([]int{1, 8, 3, 5}))
	fmt.Println(connectSticks([]int{3354, 4316, 3259, 4904, 4598, 474, 3166, 6322, 8080, 9009}))
}

func connectSticks(sticks []int) int {

	min := &minHeap{}
	heap.Init(min)
	for _, n := range sticks {
		heap.Push(min, n)
	}

	ans := 0
	for min.Len() > 1 {
		one := heap.Pop(min).(int)
		two := heap.Pop(min).(int)
		ans += one + two
		heap.Push(min, one+two)
	}

	return ans
}

type minHeap []int

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
