package main

import "fmt"

func main() {
	// Loucamente, o while de Go é um for sem ';'

	var sum int = 0

	for sum < 10 {
		sum += 1
	}

	fmt.Println(sum)
}
