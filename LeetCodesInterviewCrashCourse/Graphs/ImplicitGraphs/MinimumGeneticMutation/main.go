package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	fmt.Println(minMutation("AACCTTGG", "AATTCCGG", []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}))
}

type Pair struct {
	node  string
	steps int
}

func minMutation(startGene string, endGene string, bank []string) int {

	if len(bank) == 0 {
		return -1
	}

	if startGene == endGene {
		return 0
	}

	seen := make(map[string]bool)
	queue := list.New()

	queue.PushBack(Pair{
		node:  startGene,
		steps: 0,
	})

	for queue.Len() > 0 {

		front := queue.Front()
		queue.Remove(front)

		pair := front.Value.(Pair)

		if pair.node == endGene {
			return pair.steps
		}

		for _, neighbor := range neighbors(pair.node) {
			if ok := seen[neighbor]; !ok && inBank(bank, neighbor) {
				seen[neighbor] = true

				queue.PushBack(Pair{
					node:  neighbor,
					steps: pair.steps + 1,
				})
			}
		}
	}

	return -1
}

func neighbors(node string) []string {

	result := make([]string, 0)

	for i := 0; i < len(node); i++ {
		for _, l := range []uint8{'A', 'C', 'G', 'T'} {
			neighbor := []uint8(node)
			neighbor[i] = l
			result = append(result, string(neighbor))
		}
	}

	return result
}

func inBank(bank []string, gen string) bool {

	for _, s := range bank {
		if s == gen {
			return true
		}
	}

	return false
}
