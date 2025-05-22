package main

import "fmt"

func quickSort(v []int, ini, fim int) {
	if ini < fim {
		iPivot := partition(v, ini, fim)
		quickSort(v, ini, iPivot-1)
		quickSort(v, iPivot+1, fim)
	}
}

func partition(v []int, ini, fim int) int {
	iPivot := fim
	pIndex := ini
    for i := ini; i < iPivot; i++ {
		if v[i] <= v[iPivot] {
			v[i], v[pIndex] = v[pIndex], v[i]
			pIndex++
		}
	}
	v[iPivot], v[pIndex] = v[pIndex], v[iPivot]
	return pIndex
}

func main() {
    v := []int{4,3,2,1}
    quickSort(v, 0, 3)
    fmt.Println(v)
}
