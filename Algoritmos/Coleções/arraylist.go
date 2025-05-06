package main

import "fmt"

type List interface {
	Size() int
	Get(index int) int
	Add(e int)
	AddOnIndex(e int, index int)
	Remove(index int)
}

type ArrayList struct {
	v        []int
	inserted int
}

func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}

func (l *ArrayList) Get(index int) int {
	if index >= 0 && index < l.inserted {
		return l.v[index]
	} else {
		return -1 // erro
	}
}

func (l *ArrayList) Size() int {
	return l.inserted
}

func (l *ArrayList) doubleV() {
	newV := make([]int, 2*l.inserted)
	for i := 0; i < l.inserted; i++ {
		newV[i] = l.v[i]
	}

	l.v = newV
}

func (l *ArrayList) Add(e int) {
	if l.inserted == len(l.v) {
		l.doubleV()
	}
	l.v[l.inserted] = e
	l.inserted++
}

func (l *ArrayList) AddOnIndex(e, index int) {
	result := make([]int, l.inserted+1)
	listaTemp1 := make([]int, index+1)

	for i := 0; i < (l.inserted-(l.inserted-index))+1; i++ {
		if i == (l.inserted - (l.inserted - index)) {
			listaTemp1[i] = e
		} else {
			listaTemp1[i] = l.Get(i)
		}
	}

	listaTemp2 := make([]int, (l.inserted - index))

	indexListaTemp2 := 0

	for i := index; i <= l.inserted-1; i++ {
		listaTemp2[indexListaTemp2] = l.Get(i)
		indexListaTemp2++
	}

	indexListaTemp2 = 0

	for i := 0; i < l.inserted+1; i++ {
		if i <= index {
			result[i] = listaTemp1[i]
		} else {
			result[i] = listaTemp2[indexListaTemp2]
			indexListaTemp2++
		}
	}

	l.inserted++
	l.v = result
}

// 0 1 2 3 4

func (l *ArrayList) Remove(index int) {
	result := make([]int, l.inserted-1)
	listaTemp1 := make([]int, index)

	for i := 0; i < (l.inserted - (l.inserted - index)); i++ {
		listaTemp1[i] = l.Get(i)
	}

	listaTemp2 := make([]int, (l.inserted - index - 1))

	indexListaTemp2 := 0

	for i := index + 1; i <= l.inserted-1; i++ {
		listaTemp2[indexListaTemp2] = l.Get(i)
		indexListaTemp2++
	}

	indexListaTemp2 = 0

	for i := 0; i < l.inserted-1; i++ {
		if i < index {
			result[i] = listaTemp1[i]
		} else {
			result[i] = listaTemp2[indexListaTemp2]
			indexListaTemp2++
		}
	}

	l.inserted--
	l.v = result
}

func (l *List) Reverse() {
    for i := range l.inserted / 2 {
        l.v[i], l.v[l.inserted-i-1] = l.v[l.inserted-i-1], l.v[i]
    }
}

func main() {
	l := &ArrayList{}
	l.Init(5)
	for i := 0; i < 7; i++ {
		l.Add(i)
	}

	l.Remove(1)

	for i := 0; i <= l.inserted-1; i++ {
		fmt.Println(l.Get(i))
	}
}
