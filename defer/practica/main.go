package main

import (
	"fmt"
	"os"
)

// En la practica defer sirve para:
// limpiar recursos, cerrar archivos,
// cerrar conexiones de red, cerrar controladores de bases de datos

func main() {
	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Printf("Ocurrio un error al crear el archivo: %v", err)
		return
	}

	// Una opcion viable es usar defer para que la funcion de cerrar Close()
	// se ejecute al final, y se pone antes de terminar para que no lo olvidemos al final
	defer file.Close()

	// la funcion Write devuelve un int que es el numero
	// de bytes que fueron escritos y un err
	// en este caso solo nos interesa e err
	_, err = file.Write([]byte("Hello Gophers"))
	if err != nil {
		// si usamos defer despues de crear el archivo en la funcion cerrar archivo Close()
		// podemos omitir usarla aca file.Close()
		// file.Close()
		fmt.Printf("Ocurrio un error al escribir el archivo: %v", err)
		return
	}

	// Podemos usar la funcion Close() al final para cerrar el archivo
	// file.Close()
}
