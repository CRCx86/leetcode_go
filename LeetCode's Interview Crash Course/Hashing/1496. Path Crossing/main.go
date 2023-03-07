package main

import "fmt"

func main() {

	fmt.Println(isPathCrossing("NES"))
	fmt.Println(isPathCrossing("NESWW"))
	fmt.Println(isPathCrossing("SS"))
	fmt.Println(isPathCrossing("SN"))

}

func isPathCrossing(path string) bool {

	x, y := 0, 0
	counts := make(map[string]bool)
	counts[fmt.Sprintf("%d%d", x, y)] = true
	for i := range path {
		if path[i] == 'N' {
			y++
		}
		if path[i] == 'S' {
			y--
		}
		if path[i] == 'E' {
			x++
		}
		if path[i] == 'W' {
			x--
		}
		was := fmt.Sprintf("%d%d", x, y)
		if counts[was] {
			return true
		}
		counts[was] = true
	}

	return false

}
