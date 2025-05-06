package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	peak     *Node
	inserted int
}

func (s *Stack) Push(e int) {
	newNode := &Node{next: s.peak, value: e}
	s.peak = newNode
}

func (s *Stack) Get() int {
	return s.peak.value
}

func (s *Stack) Pop() {
	//peak := s.peak.value // Você pode retornar o valor do antigo peak
	s.peak = s.peak.next
}

func main() {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Pop()
	fmt.Println(s.Get())
}
