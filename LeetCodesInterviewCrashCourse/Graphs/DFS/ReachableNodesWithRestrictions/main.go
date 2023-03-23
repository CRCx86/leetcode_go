package main

import "fmt"

func main() {
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5}))
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}}, []int{4, 2, 1}))
	fmt.Println(reachableNodes(10, [][]int{{8, 2}, {2, 5}, {5, 0}, {2, 7}, {1, 7}, {3, 8}, {0, 4}, {3, 9}, {1, 6}}, []int{9, 8, 4, 5, 3, 1}))
}

var seen []bool
var graph map[int][]int
var r map[int]bool

func reachableNodes(n int, edges [][]int, restricted []int) int {

	seen = make([]bool, n)

	r = make(map[int]bool)
	for _, v := range restricted {
		r[v] = true
	}

	graph = make(map[int][]int)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	ans := 1
	for i := 0; i < n; i++ {
		if !seen[i] {
			seen[i] = true
			ans += dfs(i, r[i])
		}
	}

	return ans

}

func dfs(node int, isRestricted bool) int {
	count := 0
	for _, i := range graph[node] {
		if !seen[i] {
			seen[i] = true
			if !r[i] && !isRestricted {
				count++
			}
			count += dfs(i, r[i] || isRestricted)
		}
	}
	return count
}
