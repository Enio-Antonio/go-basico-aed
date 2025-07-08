package main

import "fmt"

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

//func insertionSortS(v []int) {
//    for i := 1; i < len(v); i++ {
//        c := i
//        for ; c > 0 && v[c-1] > v[c]; c-- {
//            v[c], v[c-1] = v[c-1], v[c]
//            fmt.Println(v)
//        }
//    }
//}


func main() {
	v := []int{4, 3, 2, 1, 4, 2}
	insertionSort(v)
	fmt.Println(v)
}
