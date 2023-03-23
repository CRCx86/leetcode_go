package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(nearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}))
	fmt.Println(nearestExit([][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}))
	fmt.Println(nearestExit([][]byte{{'.', '+'}}, []int{0, 0}))
}

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func nearestExit(maze [][]byte, entrance []int) int {

	m := len(maze)
	n := len(maze[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[entrance[0]][entrance[1]] = true

	queue := list.New()
	queue.PushBack([]int{entrance[0], entrance[1], 0}) //row, col, steps

	hasStep := false
	for queue.Len() > 0 {
		curr := queue.Front()
		queue.Remove(curr)

		row, col, steps := curr.Value.([]int)[0], curr.Value.([]int)[1], curr.Value.([]int)[2]
		if hasStep && (row == 0 || col == 0 || row == m-1 || col == n-1) {
			return steps
		}

		hasStep = true
		for _, direct := range directions {
			nextRow, nextCol := row+direct[0], col+direct[1]
			if valid(nextRow, nextCol, n, m, maze) && !seen[nextRow][nextCol] {
				seen[nextRow][nextCol] = true
				queue.PushBack([]int{nextRow, nextCol, steps + 1})
			}
		}
	}

	return -1
}

func valid(row int, col int, n int, m int, maze [][]byte) bool {
	return 0 <= row && row < m && 0 <= col && col < n && maze[row][col] == '.'
}
