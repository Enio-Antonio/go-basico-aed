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

func (l *LinkedList) AddOnIndex(e, index int) {
	newNode := &node{val: e, next: nil}

	if index == 0 {
		newNode.next = l.head
		l.head = newNode
	}

	aux := l.head

	for range index - 1 {
		aux = aux.next
	}

	newNode.next = aux.next
	aux.next = newNode

	l.inserted++
}

// 0 1 2 3 4 5
//   a   n

func (l *LinkedList) Remove(index int) {
	if index == 0 {
		l.head = l.head.next
		l.inserted--
	}

	aux := l.head

	for range index - 1 {
		aux = aux.next
	}

	// aux.next aponta para o 4, tem que apontar para o 5
	noaux := aux
	aux = aux.next // aqui ele é o quatro e aponta para o 5
	noaux.next = aux.next

	l.inserted--
}

// 0 1 2 3 4 5
//     a n

func (l *List) Reverse() {
    current := l.head
    var sucessor *Node
    var antecessor *Node
    for range l.inserted {
        sucessor = current.next
        current.next = antecessor
        antecessor = current
        current = sucessor
    }
    l.head = antecessor
}

func main() {
	ll := &LinkedList{}
	ll.Add(1)
	ll.Add(2)
	ll.AddOnIndex(3, 2)
	ll.Remove(2)
	for i := range ll.inserted {
		fmt.Println(ll.Get(i))
	}
}
