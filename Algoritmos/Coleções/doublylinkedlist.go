package main

import "fmt"

type Node struct {
	value    int
	next     *Node
	previous *Node
}

type DoubleLinkedList struct {
	inserted int
	head     *Node
	tail     *Node
}

func (dll *DoubleLinkedList) Add(e int) { // Totalmente errado kkkkkkkkkkkkk
	newNode := &Node{value: e, next: nil, previous: nil}
	if dll.inserted == 0 {
		dll.head = newNode
		dll.tail = newNode
		dll.inserted++
	} else {
		aux := dll.head
		for aux.next != nil {
			aux = aux.next
		}
		//newNode.previous = aux
		//aux.next = newNode
		//aux = newNode
		//dll.inserted++

    dll.tail.next = newNode
    dll.tail = newNode
	}
}

func (dll *DoubleLinkedList) Get(index int) int {
	aux := dll.head
	for range index {
		aux = aux.next
	}
	return aux.value
}

func (dll *DoubleLinkedList) AddOnIndex(e, index int) {
	if index < 0 || index > dll.inserted {
		panic("Deu ruim")
	} else {
		newNode := &Node{value: e}
		if dll.inserted == 0 {
			dll.head = newNode
			dll.tail = newNode
		} else {
			if index == 0 {
				newNode.next = dll.head
				dll.head.previous = newNode
				dll.head = newNode
			} else if index == dll.inserted {
				newNode.previous = dll.tail
				dll.tail.next = newNode
				dll.tail = newNode
			} else {
        aux := dll.head
        for range index -1 {
          aux = aux.next
        }
        newNode.previous = aux
        newNode.next = aux.next
        aux.next = newNode
        newNode.next.previous = newNode // Essa linha aqui tá certa? Não sei...
			}
		}
	}
}

func main() {
	dll := &DoubleLinkedList{}
	dll.Add(1)
	dll.Add(2)
	dll.Add(3)
	fmt.Println(dll.Get(0))
	fmt.Println(dll.Get(1))
	fmt.Println(dll.Get(2))
}
