package main

import "fmt"

func main() {
	var food [3]string
	food[0] = "🍔"
	food[1] = "🍕"
	food[2] = "🌭"
	// Esta asignacion genera un error, porque en Go los arrays se definen con un tamaño fijo
	//food[3] = "🍟"

	// Arrays literales
	food2 := [3]string{"🍔", "🍕", "🌭"}

	fmt.Println(food)
	fmt.Println(food2)
}
