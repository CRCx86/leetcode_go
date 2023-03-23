package main

import "fmt"

func main() {
	fmt.Println(countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}))
	fmt.Println(countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
}

var seen []bool
var graph map[int][]int

func countComponents(n int, edges [][]int) int {

	seen = make([]bool, n)
	graph = make(map[int][]int)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	ans := 0
	for i := 0; i < n; i++ {
		if !seen[i] {
			ans++
			seen[i] = true
			dfs(i)
		}
	}

	return ans
}

func dfs(node int) {
	for _, i := range graph[node] {
		if !seen[i] {
			seen[i] = true
			dfs(i)
		}
	}
}
