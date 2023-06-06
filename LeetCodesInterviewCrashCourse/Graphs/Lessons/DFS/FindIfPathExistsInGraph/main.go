package main

import "fmt"

var seen []bool
var graph map[int][]int

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
}

func validPath(n int, edges [][]int, source int, destination int) bool {

	seen = make([]bool, n)
	graph = make(map[int][]int)

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	seen[source] = true
	return dfs(source, destination)
}

func dfs(node int, destination int) bool {
	for _, direct := range graph[node] {
		if !seen[direct] {
			seen[direct] = true
			dfs(direct, destination)
		}
	}
	return seen[destination]
}
