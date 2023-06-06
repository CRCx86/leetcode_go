package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minStoneSum([]int{5, 4, 9}, 2))
	fmt.Println(minStoneSum([]int{4, 3, 6, 7}, 3))
}

func minStoneSum(piles []int, k int) int {

	h := &IntHeap{}
	heap.Init(h)
	for _, p := range piles {
		heap.Push(h, p)
	}

	for k > 0 {
		f := math.Round(float64(heap.Pop(h).(int)) / 2)
		heap.Push(h, int(f))
		k--
	}

	ans := 0
	for h.Len() > 0 {
		ans += heap.Pop(h).(int)
	}

	return ans
}

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
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
