package main

import "fmt"

// Por enquanto vai ser out-of-place pq tá difícil

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	inserted int
	head     *Node
}

func (l *LinkedList) add(e int) {
	newNode := Node{value: e, next: nil}
	if l.inserted == 0 {
		l.head = &newNode
	} else {
		aux := l.head
		for aux.next != nil {
			aux = aux.next
		}
		aux.next = &newNode
	}
	l.inserted++
}

func (l *LinkedList) get(index int) int {
	if l.inserted == 0 {
		return -1
	}
	aux := l.head
	for range index {
		aux = aux.next
	}
	return aux.value
}

func (l *LinkedList) remove(index int) {
	aux := l.head
	if index == 0 {
		antigo := l.head
		l.head = antigo.next
	} else {
		for range index - 1 {
			aux = aux.next
		}
		antigo := aux.next
		aux.next = antigo.next
	}
	l.inserted--
}

func selectionSort(v LinkedList) LinkedList {
	var noval LinkedList
	var index_aux int
	menor := v.get(0)
	vbackup := v

	for range v.inserted {
		for i := range vbackup.inserted {
			fmt.Printf("%d ", vbackup.get(i))
		}
		fmt.Println()
		for c := range vbackup.inserted {
			if vbackup.get(c) <= menor {
				menor = vbackup.get(c)
				index_aux = c
			}
		}
		noval.add(menor)
		vbackup.remove(index_aux)
		menor = vbackup.get(0)
	}

	return noval
}

func main() {
	var vetor LinkedList
	vetor.add(1)
	vetor.add(2)
	vetor.add(5)
	vetor.add(4)
	vetor.add(2)
	vetor.add(3)
	result := selectionSort(vetor)
	fmt.Println("Resultado:")
	for i := range result.inserted {
		fmt.Printf("%d ", result.get(i))
	}
}
