package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(predictPartyVictory("RD"))    // R
	fmt.Println(predictPartyVictory("RDD"))   // D
	fmt.Println(predictPartyVictory("DDRRR")) // D
}

// ключ: сравниваем по индексам, меньший индекс побеждает и переезжает в конец очереди, становясь большим
// следующему отличному символу он уже проиграет
func predictPartyVictory(senate string) string {

	R := list.New()
	D := list.New()

	n := len(senate)

	for i := range senate {
		if senate[i] == 'R' {
			R.PushBack(i)
		} else {
			D.PushBack(i)
		}
	}

	for R.Len() > 0 && D.Len() > 0 {

		r := R.Front().Value.(int)
		R.Remove(R.Front())
		d := D.Front().Value.(int)
		D.Remove(D.Front())

		if r < d {
			R.PushBack(r + n)
		} else {
			D.PushBack(d + n)
		}

	}

	if R.Len() > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
