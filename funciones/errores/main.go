package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Go se diseño para manejar los errores en el momento
	// en el momento en que se presentan - No tiene excepciones
	content, err := ioutil.ReadFile("./hello.txt")
	// controlamos el error al verificar que la variable
	// err sea diferente de su valor cero, usando la expresion
	// if err != nil, si esto es true, ocurrio un error
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}

	// Esta parte del programa solo se ejecuta si no hubo errores
	fmt.Println(string(content))

}
