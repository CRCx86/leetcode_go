package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 1))
	fmt.Println(topKFrequent([]string{"a", "aa", "aaa"}, 3))
}

func topKFrequent(words []string, k int) []string {

	counts := make(map[string]int)
	for _, s := range words {
		counts[s]++
	}

	h := &intHeap{}
	heap.Init(h)
	for key, val := range counts {
		heap.Push(h, &pair{
			key:   val,
			value: key,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]string, h.Len())
	i := h.Len() - 1
	for h.Len() > 0 {
		ans[i] = heap.Pop(h).(*pair).value
		i--
	}

	return ans
}

type pair struct {
	key   int
	value string
}

type intHeap []*pair

func (h *intHeap) Len() int { return len(*h) }
func (h *intHeap) Less(i, j int) bool {
	if (*h)[i].key != (*h)[j].key {
		return (*h)[i].key < (*h)[j].key
	}
	return (*h)[i].value > (*h)[j].value
}
func (h *intHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
