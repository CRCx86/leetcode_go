package main

import "fmt"

func main() {
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
	fmt.Println(destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}))
	fmt.Println(destCity([][]string{{"A", "Z"}}))
}

func destCity(paths [][]string) string {

	// неоптимальное
	//source := make(map[string]int)
	//dest := make(map[string]int)
	//for i := range paths {
	//	source[paths[i][0]]++
	//	dest[paths[i][1]]++
	//}
	//
	//for k, _ := range dest {
	//	if _, ok := source[k]; ok {
	//		dest[k]--
	//	}
	//}
	//
	//for k, v := range dest {
	//	if v == 1 {
	//		return k
	//	}
	//}

	// оптимальное - leetcode
	dest := make(map[string]int)
	for i := range paths {
		dest[paths[i][1]]++
	}

	for i := range paths {
		delete(dest, paths[i][0])
	}

	for k := range dest {
		return k
	}

	return ""
}
