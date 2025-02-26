package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += 1
	}

	fmt.Println(sum)

	// Que louco, o "i := 0" e "i++" são opcionais
	var sum2 int = 0
	for ;sum2 < 10; {
		sum2 += 1
	}

	fmt.Println(sum2)
}
