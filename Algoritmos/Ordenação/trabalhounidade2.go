package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func bubbleSort(v []int) {
	for n := range len(v) - 1 {
		trocou := false
		for i := range len(v) - 1 - n {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				trocou = true
			}
		}
		if !trocou {
			return
		}
	}
}

func insertionSort(v []int) {
	for insercao := 1; insercao < len(v); insercao++ {
		temp := v[insercao]
		i := insercao
		for ; i > 0 && v[i-1] > temp; i-- {
			//v[i-1], v[i] = v[i], v[i-1]
			v[i] = v[i-1]
		}
		v[i] = temp
	}
}

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

func createV(n int, sorted bool) []int {
	v := make([]int, n)
	for i := range n {
		v[i] = rand.Intn(100)
	}
	if sorted {
		bubbleSort(v)
	}
	return v
}

func main() {
	v := createV(10000, false)
	vis := v
	vbs := v
	vms := v
	vqs := v

	start := time.Now()

	// Code to be measured
	_ = selectionSort(v)

	elapsed := time.Since(start)
	fmt.Printf("Execution time selec: %s\n", elapsed)

	start = time.Now()

	// Code to be measured
	bubbleSort(vbs)

	elapsed = time.Since(start)
	fmt.Printf("Execution time bubble: %s\n", elapsed)

	start = time.Now()

	// Code to be measured
	mergeSort(vms)

	elapsed = time.Since(start)
	fmt.Printf("Execution time merge: %s\n", elapsed)

	start = time.Now()

	// Code to be measured
	quickSort(vqs, 0, len(vqs)-1)

	elapsed = time.Since(start)
	fmt.Printf("Execution time quick: %s\n", elapsed)

	start = time.Now()

	// Code to be measured
	insertionSort(vis)

	elapsed = time.Since(start)
	fmt.Printf("Execution time insertion: %s\n", elapsed)
}
