package main

import "fmt"

func bubbleSort(v []int) {
    var trocou bool
    for n := range len(v)-1 {
        trocou = false
        for i := range len(v) - 1 - n {
            if (v[i] > v[i+1]) {
                v[i], v[i+1] = v[i+1], v[i]
                trocou = true
            }
        } 
        if !trocou {return}
    }
}

func main() {
    vetor := []int{1,2,3,4,5}
    bubbleSort(vetor)
    fmt.Println(vetor)
}
