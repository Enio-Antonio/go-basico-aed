package main

import "fmt"

func main() {
	var x int = 10
	var y int = 10

	// If normal
	if x == y {
		fmt.Println("Passou no teste.")
	}

	// If com else
	if x != y {
		fmt.Println("Passou no teste.")
	} else {
		fmt.Println("Passou no if.")
	}
}
