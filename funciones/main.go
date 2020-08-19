package main

import "fmt"

func main() {
	hello()
	hello1("Jhon", "Guerrero")
	emoji := "🐶"
	change(emoji)
	fmt.Println("Por valor: ", emoji)
	change1(&emoji)
	fmt.Println("Por referencia: ", emoji)
}

// funciones sin parametros
func hello() {
	fmt.Println("Hello EDteam")
}

// funciones con parametros
func hello1(firstName string, lastName string) {
	fmt.Printf("Hello %s %s\n", firstName, lastName)
}

// funcion que recibe una copia del valor, no cambnia el emoji original 🐶
func change(value string) {
	value = "😿"
}

// recibe un puntero a la variable emoji que contiene 🐶
// Accede por medio del operador * y cambia al emiji 😿
func change1(value *string) {
	*value = "😿"
}
