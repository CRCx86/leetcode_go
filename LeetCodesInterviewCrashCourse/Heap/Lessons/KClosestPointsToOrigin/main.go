package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}

func kClosest(points [][]int, k int) [][]int {

	h := &pairs{}
	heap.Init(h)

	for _, col := range points {
		x := col[0]
		y := col[1]
		heap.Push(h, &pair{
			x:        x,
			y:        y,
			distance: math.Sqrt(math.Pow(float64(x)-0, 2) + math.Pow(float64(y)-0, 2)),
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([][]int, h.Len())
	i := 0
	for h.Len() > 0 {
		p := heap.Pop(h).(*pair)
		ans[i] = make([]int, 2)
		ans[i][0] = p.x
		ans[i][1] = p.y
		i++
	}

	return ans
}

type pair struct {
	x        int
	y        int
	distance float64
}

type pairs []*pair

func (h *pairs) Len() int { return len(*h) }
func (h *pairs) Less(i, j int) bool {
	return (*h)[i].distance > (*h)[j].distance
}
func (h *pairs) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *pairs) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}
func (h *pairs) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
