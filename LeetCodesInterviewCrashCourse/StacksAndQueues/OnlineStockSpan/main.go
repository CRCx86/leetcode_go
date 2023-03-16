package main

import (
	"container/list"
	"fmt"
)

func main() {
	constructor := Constructor()
	fmt.Println(constructor.Next(100))
	fmt.Println(constructor.Next(80))
	fmt.Println(constructor.Next(60))
	fmt.Println(constructor.Next(70))
	fmt.Println(constructor.Next(60))
	fmt.Println(constructor.Next(75))
	fmt.Println(constructor.Next(85))

	fmt.Println(constructor.Next2(100))
	fmt.Println(constructor.Next2(80))
	fmt.Println(constructor.Next2(60))
	fmt.Println(constructor.Next2(70))
	fmt.Println(constructor.Next2(60))
	fmt.Println(constructor.Next2(75))
	fmt.Println(constructor.Next2(85))
}

type StockSpanner struct {
	in *list.List
	//out *list.List
}

func Constructor() StockSpanner {
	return StockSpanner{
		in: list.New(),
		//out: list.New(),
	}
}

// Next without stack
func (t *StockSpanner) Next(price int) int {

	ans := 1

	for e := t.in.Back(); e != nil; e = e.Prev() {
		i := e.Value.(int)
		if i > price {
			break
		}
		ans++
	}

	t.in.PushBack(price)

	return ans

}

// Next2 with stack (leetcode): идея в том, чтобы хранить последнее суммарное значение по дням в паре с ценой
// Если цена на следующий день будет выше - то суммируем с последним вычисленным значением
// Если цена будет ниже - будет всё равно 1
func (t *StockSpanner) Next2(price int) int {

	ans := 1

	for t.in.Len() > 0 && t.in.Back().Value.([2]int)[0] <= price {
		ans += t.in.Back().Value.([2]int)[1]
		t.in.Remove(t.in.Back())
	}

	e := [2]int{price, ans}
	t.in.PushBack(e)

	return ans

}
