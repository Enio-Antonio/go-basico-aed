package main

import "fmt"

type Stack struct {
	inserted int
	v        []int
}

func (s *Stack) Init(size int) {
	s.v = make([]int, 5)
}

func (s *Stack) Push(e int) {
	s.v[s.inserted] = e
	s.inserted++
}

func (s *Stack) Peak() int {
	return s.v[s.inserted-1]
}

func (s *Stack) Pop() {
	if s.inserted == 0 {
		panic("A lista tá vazia, como é que vc quer remover um elemento?")
	} else {
		s.inserted--
	}
}

func main() {
	//s := &Stack{inserted: 0, v: nil}
	var s Stack
	s.Init(5)
	s.Push(1)
	s.Push(2)
	s.Pop()
	fmt.Println(s.Peak())
}
