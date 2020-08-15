package main

import "fmt"

func main() {
	// operadores aritmeticos () * / % + -
	var a = 4 + 3*5
	fmt.Println(a)

	// Operadores de asignacion: = += -= *= /= %=
	var b = 10
	b += 2
	fmt.Println(b)

	// Declaracion post-incremento y post-decremento: ++ --
	// no son una expresion sino una declaracion

	var c = 3
	// Genera un error porque se esta utilzando como expresion
	// c = c++ + 2

	c--

	fmt.Println(c)

	// Operadores de comparacion > < == != >= <=
	fmt.Println(4 > 5)
	fmt.Println(4 == 4)
	fmt.Println(4 != 4)

	// Operadores logicos && ||
	var age = 30
	fmt.Println("Edad ", age >= 18 && age <= 60)
	fmt.Println("niño o viejo: ", age < 18 || age > 60)

	// Operador Unario !
	fmt.Println(!(4 == 4))
}
