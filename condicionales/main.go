package main

import "fmt"

func main() {
	emoji := "🦄"

	if emoji == "🌵" {
		fmt.Println("Es un cactus")
	} else if emoji == "🙂" {
		fmt.Println("Es una cvarita")
	} else {
		fmt.Printf("El emoji es: %s", emoji)
	}

	fmt.Println("\n=================================================")

	// Esta forma del if me permite usar la variable emoji1 solo dentro del condicional
	if emoji1 := "🙀"; emoji1 == "🌵" {
		fmt.Println("Es un cactus")
	} else if emoji1 == "🙂" {
		fmt.Println("Es una cvarita")
	} else {
		fmt.Printf("El emoji es: %s", emoji1)
	}

	// No puedo utilizar la variable emoji1 fuera del condional
	//	fmt.Println(emoji1)
}
