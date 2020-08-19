package main

import "fmt"

// defer es diferir o aplazar la ejecucion de una funcion
func main() {

	// defer proboca que la impresion del numero 1 se haga al final
	// el defer provoca que se agregue a una pila
	// para este ejemplo se va agregando de arriba hacia abajo
	// y se ejecutan de la ultima a la primera, comportamiento
	// normal de las structuras pilas
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)

	a := 5
	// al usar defer se aplaza la ejecucion de la funcion, pero los
	// argumentos son evluados inmediatamente
	defer fmt.Println("Defer: ", a)

	a = 10
	fmt.Println(a)
}
