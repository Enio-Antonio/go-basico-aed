package main

import "fmt"

func main() {
	var vetor = [2]int {1, 2}

	// Outra forma de inicializar
	var vetor_2 [2]int 
	for i := 0; i < 2; i++ {
		vetor_2[i] = i
	}

	// Percorrendo normal
	for i := 0; i < len(vetor); i++ {
		fmt.Println(vetor[i])
	}

	// Percorrendo com uma variável
	for _, elemento := range vetor { // Não for usar o índice, o nome deve ser '_'
		fmt.Println(elemento)
	}

	// Ignorar (se não acontece esse de compilação)
	fmt.Println(vetor_2[0])
	
}
