package main

import(
    "fmt"
    "math"
)

func selectionSortOP(v []int) []int {
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

func selectionSort(v []int) {
    for varredura := 0; varredura < len(v)-1; varredura++ {
        iMenor := varredura
        for i := varredura + 1; i < len(v); i++ {
            if v[i] < v[iMenor] {
                iMenor = i                                 
            }                 
        }
        v[varredura], v[iMenor] = v[iMenor], v[varredura]                     
    } 
}

func main() {
    vetor := []int{4,3,2,1,1}
    selectionSort(vetor)
    fmt.Println(vetor)
}
