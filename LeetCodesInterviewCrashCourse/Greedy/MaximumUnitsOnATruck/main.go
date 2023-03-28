package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {

	h := &pairHeap{}
	heap.Init(h)

	for i := 0; i < len(boxTypes); i++ {
		heap.Push(h, &pair{
			box:   boxTypes[i][0],
			units: boxTypes[i][1],
		})
	}

	ans := 0
	for h.Len() > 0 {
		p := heap.Pop(h).(*pair)
		truckSize -= p.box
		ans += p.box * p.units
		if truckSize == 0 {
			return ans
		} else if truckSize < 0 {
			ans += truckSize * p.units
			return ans
		}
	}

	return ans
}

type pair struct {
	box   int
	units int
}

type pairHeap []*pair

func (h *pairHeap) Len() int { return len(*h) }
func (h *pairHeap) Less(i, j int) bool {
	return (*h)[i].units > (*h)[j].units
}
func (h *pairHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *pairHeap) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}
func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
