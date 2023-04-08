package main

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {}, {4}, {}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {

	ans := make([][]int, 0)
	curr := make([]int, 1)
	backtrack(curr, &ans, graph, 0)
	return ans
}

func backtrack(curr []int, ans *[][]int, graph [][]int, node int) {

	if node == len(graph)-1 {
		dist := make([]int, len(curr))
		copy(dist, curr)
		*ans = append(*ans, dist)
	}

	for _, nextNode := range graph[node] {
		if find(curr, nextNode) == -1 {
			curr = append(curr, nextNode)
			backtrack(curr, ans, graph, nextNode)
			curr = curr[0 : len(curr)-1]
		}
	}
}

//func find[T comparable](slice []T, x T) int {
//
//	for i := range slice {
//		if slice[i] == x {
//			return i
//		}
//	}
//
//	return -1
//}

func find(slice []int, x int) int {

	for i := range slice {
		if slice[i] == x {
			return i
		}
	}

	return -1
}
