package main

import "fmt"

type List interface {
	Size() int
	Get(index int) int
	Add(e int)
	AddOnIndex(e, index int)
	Remove(index int)
}

type LinkedList struct {
	head     *node
	inserted int
}
a
type node struct {
	val  int
	next *node
}

func (l *LinkedList) Add(e int) {
	newNode := &node{val: e, next: nil}
	if l.inserted == 0 {
		l.head = newNode
	} else {
		aux := l.head
		for aux.next != nil {
			aux = aux.next
		}
		aux.next = newNode
	}
	l.inserted++
}

func (l *LinkedList) Size() int {
	return l.inserted
}

func (l *LinkedList) Get(index int) int {
	aux := l.head
	for range index {
		aux = aux.next
	}
	return aux.val
}

func main() {
	ll := &LinkedList{}
	ll.Add(1)
	ll.Add(2)
	valor := ll.Get(0)
	fmt.Println(valor)
}
