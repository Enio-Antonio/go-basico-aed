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
    v []int
    inserted int
}

func (l *ArrayList) Init(size int) {
    l.v = make([]int, size)
}

func (l *ArrayList) Get() int {
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
    newV := make([]int, 2 * l.inserted)
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

func main() {
    l := &ArrayList{}
    l.Init(10)
    l.Add(1)

    fmt.Println(l.Get(0))
}
