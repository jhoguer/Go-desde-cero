package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	// Go se diseño para manejar los errores en el momento
	// en el momento en que se presentan - No tiene excepciones
	content, err := ioutil.ReadFile("archivo/hello.txt")
	// controlamos el error al verificar que la variable
	// err sea diferente de su valor cero, usando la expresion
	// if err != nil, si esto es true, ocurrio un error
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}

	// Esta parte del programa solo se ejecuta si no hubo errores
	fmt.Println(string(content))

	fmt.Println("============================================")

	res, err := division(12, 2)

	if err != nil {
		fmt.Printf("Ocurrio un error al dividir: %v", err)
		return
	}

	println("El resultado de la division es: ", res)
	fmt.Println("============================================")

	fmt.Println("============================================")

	result, err := divisionNombrado(10, 0)

	if err != nil {
		fmt.Printf("Ocurrio un error al dividir: %v", err)
		return
	}

	println("El resultado de la divisionNombrada es: ", result)

}

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// para generar el error podemos usar el paquete errors
		// usamos errors.New()
		return 0, errors.New("No puedes dividir por cero")
	}

	return dividendo / divisor, nil
}

// se puede retornar valores con nombre, en este caso result y err
func divisionNombrado(dividendo, divisor int) (result int, err error) {
	if divisor == 0 {
		// para generar el error podemos usar el paquete errors
		// usamos errors.New()
		err = errors.New("No puedes dividir por cero")
		return
	}

	result = dividendo / divisor
	return
}
