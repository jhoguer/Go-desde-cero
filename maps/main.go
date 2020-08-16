package main

import "fmt"

func main() {
	animals := make(map[string]string)
	animals["cat"] = "🐈"
	animals["dog"] = "🐕"

	fmt.Println(animals)

	fruits := map[string]string{
		"apple":    "🍎",
		"banana":   "🍌",
		"pineaple": "🍍",
	}

	fmt.Println(fruits)

	// Eliminar elementos, en este caso elmento "banana"
	delete(fruits, "banana")
	fmt.Println(fruits)

	// Como el elemento "gorilla" no existe en este caso
	// imprime un string vacio
	fmt.Println(animals["gorilla"])

	// Go puede devolver mas de 1 valor
	contenido, ok := animals["gorilla"]
	fmt.Printf("Contenido:\t\t%s\nExiste:\t\t\t%v\n", contenido, ok)

	if contenido1, ok1 := animals["gorilla"]; !ok {
		animals["gorilla"] = "🦍"
		fmt.Printf("Contenido:\t\t%s\nExiste:\t\t\t%v\n", contenido1, ok1)
	}

	fmt.Println(animals)

}
