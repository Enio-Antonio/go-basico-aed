package main

import "fmt"

// Pensei que fosse um bubble sort, mas não é assim que ele funciona
// Fica aqui como curiosidade
func bubbleSort(v []int) {
    i := 0
    var aux int
    for i < len(v) - 1 { 
        if v[i] > v[i+1] {
            aux = v[i]
            v[i] = v[i+1]
            v[i+1] = aux
            i = 0
        } else {
            i++
        }
        fmt.Println(v)
    }
}

func main() {
    vetor := []int{4,3,2,1}
    bubbleSort(vetor)
    fmt.Println(vetor)
}
