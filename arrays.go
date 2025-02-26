package main

import "fmt"

func main() {
	var vetor = [2]int {1, 2}

	// Percorrendo normal
	for i := 0; i < len(vetor); i++ {
		fmt.Println(vetor[i])
	}

	// Percorrendo com uma variável
	for _, elemento := range vetor { // Não for usar o índice, o nome deve ser '_'
		fmt.Println(elemento)
	}
	
}
