package main

import "fmt"

func merge(v, e, d []int) {
    iv := 0
    ie := 0
    id := 0

    for ie < len(e) && id < len(d) {
        if e[ie] <= d[id] {
            v[iv] = e[ie]
            ie++
        } else {
            v[iv] = d[id]
            id++
        }
        iv++
    }

    for ie < len(e) {
        v[iv] = e[ie]
        ie++
        iv++
    }

    for id < len(d) {
        v[iv] = d[id]
        id++
        iv++
    }
}

func mergeSort(v []int) {
    if len(v) > 1 {
        meio := len(v) / 2
        esq := make([]int, meio)
        dir := make([]int, len(v)-meio)

        for i := range len(esq) {
            esq[i] = v[i]
        }
        for i := range len(dir) {
            dir[i] = v[i+len(esq)]
        }
        mergeSort(esq)
        mergeSort(dir)
        merge(v, esq, dir)
    }
}

func main() {
    v := []int{5,4,3,2,1}
    mergeSort(v)
    fmt.Println(v)
}
