package main

import(
    "fmt"
    "math"
)

func selectionSort(v []int) []int {
    newv := make([]int, len(v))
    menor := math.MaxInt
    aux := 0
    for i := range v {
        for c := range v {
            if v[c] < menor {
                menor = v[c]
                aux = c
            }
        }
        newv[i] = menor
        menor = math.MaxInt
        v[aux] = math.MaxInt
    }
    return newv
}

func main() {
    vetor := []int{4,3,2,1,1}
    newv := selectionSort(vetor)
    fmt.Println(newv)
}
