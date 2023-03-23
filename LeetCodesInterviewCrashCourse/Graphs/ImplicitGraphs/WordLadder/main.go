package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}

type Pair struct {
	node  string
	steps int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	if len(wordList) == 0 {
		return 0
	}

	if beginWord == endWord {
		return 0
	}

	seen := make(map[string]bool)
	queue := list.New()

	queue.PushBack(Pair{
		node:  beginWord,
		steps: 1,
	})

	bank := make(map[string]bool)
	for _, s := range wordList {
		bank[s] = true
	}

	for queue.Len() > 0 {

		front := queue.Front()
		queue.Remove(front)

		pair := front.Value.(Pair)

		if pair.node == endWord {
			return pair.steps
		}

		for _, neighbor := range neighbors(pair.node) {
			if ok := seen[neighbor]; !ok && bank[neighbor] {
				seen[neighbor] = true

				queue.PushBack(Pair{
					node:  neighbor,
					steps: pair.steps + 1,
				})
			}
		}
	}

	return 0
}

func neighbors(node string) []string {

	result := make([]string, 0)

	for i := 0; i < len(node); i++ {
		var j uint8
		for j = 0; j < 26; j++ {
			neighbor := []uint8(node)
			neighbor[i] = j + 'a'
			result = append(result, string(neighbor))
		}
	}

	return result
}
