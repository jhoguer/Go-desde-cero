package main

import "fmt"

func main() {
	// los Slices no poseen datos, son apuntadores a un Array
	set := [7]string{"🦁", "🐎", "🐄", "🦋", "🐦", "✈️", "🚆"}

	// Para generar el slice se va desde el indice del array del primer elemento
	// hasta el indice del ultimo elemento que queremos incluir + 1
	animales := set[0:5]  // "🦁", "🐎", "🐄", "🦋", "🐦"
	transport := set[5:7] // "✈️", "🚆"

	fmt.Println("Array original: ", set)

	transport[0] = "🦅"

	fmt.Println("Array Mod: ", set)
	fmt.Println("Animales: ", animales)
	fmt.Println("Transporte: ", transport)
	fmt.Println("Transporte[0]: ", transport[0])

	// Si no se especifica el primer elemeto del Array, se asume que es 0
	animales2 := set[:4] // "🦁", "🐎", "🐄", "🦋", "🐦"

	// Si no se especifica el ultimo elemeto del Array, se asume que es ultimo elemento
	transport2 := set[4:] // "✈️", "🚆"

	fmt.Println("Animales2: ", animales2)
	fmt.Println("Transporte2: ", transport2)

	// Se imprimen todos los elemento porque no se acoto a nada
	fmt.Println("All ", set[:])

	fmt.Println("______________________________________")

	food := [6]string{"🍒", "🍏", "🍍", "🥝", "🍻", "🍷"}
	fruitsGren := food[1:4] // "🍏", "🍍", "🥝"

	// Al agregar mas elementos de la capacidad que tiene el Array original food
	// se genera un nuevo Array y el Slice fruitsGren apunta a ese nuevo elemento
	// y deja de referenciar al Array original food
	fruitsGren = append(fruitsGren, "🍐", "🍉", "🍈")

	fmt.Println("Food=>", food)
	fmt.Println("Frutas Verdes=>", fruitsGren)
	fmt.Println("Cantidad de frutas=>", len(fruitsGren))
	fmt.Println("Capacidad=>", cap(fruitsGren))

	fmt.Println("=======================================")
	// fruits := []string{"🍐", "🍉", "🍈"}
	fruits := make([]string, 0, 3)

	fruits = append(fruits, "🍍", "🍒", "🍎", "🍓", "🍊", "🍐", "🍉", "🍈")

	fmt.Println("Frutas =>", fruits)
	fmt.Println("Cantidad de frutas=>", len(fruits))
	fmt.Println("Capacidad=>", cap(fruits))
	fmt.Println("=======================================")

	fmt.Println("=======================================")

	// Retorna true
	var fruits2 []string
	fmt.Println(fruits2 == nil)
	fmt.Println("Frutas =>", fruits2)
	fmt.Println("Cantidad de frutas=>", len(fruits2))
	fmt.Println("Capacidad=>", cap(fruits2))
	fmt.Println("=======================================")

	// Retorna false
	fruits3 := []string{}
	fmt.Println(fruits3 == nil)

}
