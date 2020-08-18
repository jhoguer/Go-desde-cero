package main

import "fmt"

func main() {
	emoji := "🍎"
	switch emoji {
	case "🐈", "🐶":
		fmt.Println("Es un animal")
	case "🍌", "🍎":
		fmt.Println("Es una fruta")
	default:
		fmt.Println("No es un animal ni una fruta")
	}

	// Usando operadores logicos
	emoji2 := "🐈"
	switch {
	case emoji2 == "🐈" || emoji2 == "🐶":
		fmt.Println("Es un animal")
	case emoji2 == "🍌" || emoji2 == "🍎":
		fmt.Println("Es una fruta")
	default:
		fmt.Println("No es un animal ni una fruta")
	}
}
