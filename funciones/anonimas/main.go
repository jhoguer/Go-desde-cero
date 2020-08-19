package main

import "fmt"

func main() {

	// Una forma
	x := func() {
		fmt.Println("Hola Jhon")
	}

	x()

	// Otra forma - funciones anonimas auto-ejecutadas
	func(text string) {
		fmt.Println("Hola ", text)
	}("Gophers")
}
