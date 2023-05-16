package main

import "fmt"

func main() {
	fmt.Println(robotWithString("zza"))      // "azz"
	fmt.Println(robotWithString("bac"))      // "abc"
	fmt.Println(robotWithString("bdda"))     // "addb"
	fmt.Println(robotWithString("bydizfve")) // "bdevfziy"
}

func robotWithString(s string) string {

	// ключевая часть решения: отслеживанием текущий минимальный символ в строке
	currMin := func(arr []uint32) uint8 {
		for i := 0; i < 26; i++ {
			if arr[i] != 0 {
				return uint8('a' + i)
			}
		}
		return 'a'
	}

	minFreq := make([]uint32, 26)
	for i := 0; i < len(s); i++ {
		minFreq[s[i]-'a']++ // подсчитали частоту каждого символа
	}

	ans := make([]uint8, 0)
	t := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		t = append(t, s[i])                                 // сразу добавляем в стек и делаем это пока не встретим минимальный текущий символ
		minFreq[s[i]-'a']--                                 // уменьшаем количество наличий символа
		for len(t) > 0 && t[len(t)-1] <= currMin(minFreq) { // на вершине стека оказался самый минимальный из текущих
			// например: встретили b, а следующий минимальный - e
			ans = append(ans, t[len(t)-1]) // выталкиваем всё в ответ
			t = t[:len(t)-1]
		}
	}

	for len(t) > 0 { // остатки в случае, если минимальный символ был в конце
		ans = append(ans, t[len(t)-1])
		t = t[:len(t)-1]
	}

	return string(ans)
}
