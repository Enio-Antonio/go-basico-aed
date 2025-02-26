package main

import "fmt"
import "time"

func main() {
	fmt.Println("Pula linha")
	fmt.Print("Não pula linha\n")
	fmt.Print("Hora agora: ", time.Now())
}
