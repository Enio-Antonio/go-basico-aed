// Exemplo da primeira aula do cap. 10 do canal "Aprenda Go"

package main

import "fmt"

type cliente struct {
	nome string
	idade int
	fumante bool
}

func main() {
	cliente1 := cliente{
		nome: "João",
		idade: 20,
		fumante: false, // vírgula obrigatória no último campo
	}

	// Outra forma de inicializar
	cliente2 := cliente{"Joana", 30, false}

	fmt.Println(cliente1.nome)
	fmt.Println(cliente2)
}
