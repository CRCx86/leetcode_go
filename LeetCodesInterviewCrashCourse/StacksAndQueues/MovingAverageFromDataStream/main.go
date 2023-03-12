package main

import (
	"container/list"
	"fmt"
)

func main() {
	constructor := Constructor(3)
	fmt.Println(constructor.Next(1))
	fmt.Println(constructor.Next(10))
	fmt.Println(constructor.Next(3))
	fmt.Println(constructor.Next(5))
}

type MovingAverage struct {
	size    int
	current int
	queue   *list.List
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		queue: list.New(),
	}
}

func (t *MovingAverage) Next(val int) float64 {

	t.queue.PushBack(val)
	t.current++

	if t.current > t.size {
		t.queue.Remove(t.queue.Front())
		t.current--
	}

	ans := 0
	for e := t.queue.Front(); e != nil; e = e.Next() {
		ans += e.Value.(int)
	}

	size := t.size
	if t.queue.Len() < size {
		size = t.queue.Len()
	}

	return float64(ans) / float64(size)
}
