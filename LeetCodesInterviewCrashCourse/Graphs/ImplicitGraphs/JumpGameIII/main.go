package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0))
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))
}

type Pair struct {
	node   int
	weight int
}

func canReach(arr []int, start int) bool {

	if start >= len(arr) {
		return false
	}

	seen := make(map[int]bool)
	queue := list.New()

	queue.PushBack(Pair{
		node:   start,
		weight: arr[start],
	})
	seen[start] = true

	for queue.Len() > 0 {

		front := queue.Front()
		queue.Remove(front)

		pair := front.Value.(Pair)
		if pair.weight == 0 {
			return true
		}

		for _, nextIndex := range []int{pair.node + pair.weight, pair.node - pair.weight} {
			if valid(arr, nextIndex) && !seen[nextIndex] {
				seen[nextIndex] = true
				queue.PushBack(Pair{
					node:   nextIndex,
					weight: arr[nextIndex],
				})
			}
		}

	}

	return false

}

func valid(arr []int, index int) bool {
	return 0 <= index && index < len(arr)
}
