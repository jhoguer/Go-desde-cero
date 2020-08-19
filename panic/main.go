package main

import "fmt"

// La funcion panic nos permite entrar en panico
// Permite finalizar la ejecucion del programa cuando es llamada
func main() {
	division(10, 2)
	division(30, 3)
	division(12, 0)
	division(20, 4)
}

func division(dividendo, divisor int) {

	// usando la funcion recover dentro de esta funcion anonima autoejecutable
	// podemos atrapar el panic si se llega a generar
	// lo que permite continuar con la ejecucion del programa
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperandome del panico", r)
		}
	}()
	validarDivisior(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisior(divisor int) {
	if divisor == 0 {
		panic("💥")
	}
	return
}
