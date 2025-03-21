package main
import "fmt"

func somar(x int, y int) int {
	return x + y
}

// Retorno void
func printzero() {
	fmt.Println("0")
}

// Dois retornos tipo Python
func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println(somar(1, 2))
	printzero()
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
