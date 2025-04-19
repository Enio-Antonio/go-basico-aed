package main

import "fmt"

func binarySearch(v []int, e int) int {
    initial := 0
    final := len(v) - 1
    middle := (final + initial) / 2

    for final >= initial {
        if e > v[middle] {
            initial = middle + 1
            middle = (final + initial) / 2
        } else if e < v[middle] {
            final = middle - 1
            middle = (final + initial) / 2
        } else if e == v[middle] {
            return middle
        }
    }
    return -1
}

func recursiveBinarySearch(v []int, e, final, inicial int) int {
    if inicial > final {
        return -1
    }
    middle := (final + inicial) / 2
    if e < v[middle] {
        return recursiveBinarySearch(v, e, middle-1, inicial)
    } else if e > v[middle] {
        return recursiveBinarySearch(v, e, final, middle+1)
    } else {
        return middle
    }
}

func main() {
    v := []int{1,2,3,4,5}
    result := binarySearch(v, 1)
    fmt.Println(v)
    fmt.Printf("O número 1 está no índice: %d\n", result)
}
