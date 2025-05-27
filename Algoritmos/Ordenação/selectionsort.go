package main

import(
    "fmt"
    "math"
)

func selectionSort(v []int) []int {
    newv := make([]int, len(v))
    for i := range v {
        menor := 0
        for c := 1; c < len(v); c++ {
            if v[c] < v[menor] {
                menor = c
            }
        }
        newv[i] = v[menor]
        v[menor] = math.MaxInt
    }
    return newv
}

func main() {
    vetor := []int{4,3,2,1,1}
    newv := selectionSort(vetor)
    fmt.Println(newv)
}
