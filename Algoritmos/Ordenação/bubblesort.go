package main

import "fmt"

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
    }
}

func main() {
    vetor := []int{10,9,1,7,6,9,4,3,2,1}
    bubbleSort(vetor)
    fmt.Println(vetor)
}
