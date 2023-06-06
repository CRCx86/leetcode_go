package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
}

var m int
var n int
var graph [][]int
var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var seen [][]bool

func maxAreaOfIsland(grid [][]int) int {

	graph = grid
	m = len(graph)
	n = len(graph[0])

	seen = make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	max := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if graph[row][col] == 1 && !seen[row][col] {
				seen[row][col] = true
				res := dfs(row, col) + graph[row][col]
				if res > max {
					max = res
				}
			}
		}
	}

	return max
}

func dfs(row int, col int) int {

	square := 0
	for _, direction := range directions {
		nextRow, nextCol := row+direction[0], col+direction[1]
		if valid(nextRow, nextCol) && !seen[nextRow][nextCol] {
			seen[nextRow][nextCol] = true
			square += dfs(nextRow, nextCol) + graph[nextRow][nextCol]
		}
	}
	return square
}

func valid(row int, col int) bool {
	return 0 <= row && row < m && 0 <= col && col < n && graph[row][col] == 1
}
