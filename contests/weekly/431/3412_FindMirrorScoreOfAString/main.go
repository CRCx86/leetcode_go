package main

func main() {

}

func calculateScore(s string) int64 {
	var ans int64
	m := [26][]int{}

	for i := range s {
		mirrIdx := 'z' - s[i]
		idxList := len(m[mirrIdx])

		if idxList > 0 {
			idx := m[mirrIdx][len(m[mirrIdx])-1]
			m[mirrIdx] = m[mirrIdx][:len(m[mirrIdx])-1]
			ans += int64(i - idx)
		} else {
			m[s[i]-'a'] = append(m[s[i]-'a'], i)
		}
	}

	return ans
}
