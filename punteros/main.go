package main

import "fmt"

func main() {
	fruit := "🍎"
	var p *string
	p = &fruit
	fmt.Println("_____________________________________________")
	fmt.Printf("Variable: fruit\nTipo:\t\t\t %T \nValor:\t\t\t %s \nDireccion:\t\t %v \n", fruit, fruit, &fruit)
	fmt.Println("_____________________________________________")

	fruit = "🍌"

	// El operador de desreferenciacion * se utiliza para acceder al valor que se almacena
	// en la direccion de memoria a la que esta haciendo referencia el puntero
	fmt.Printf("puntero: p\nTipo:\t\t\t %T \nValor:\t\t\t %v \nDireccion:\t\t %v \nDesreferen:\t\t %s \n", p, p, &p, *p)
	fmt.Println("_____________________________________________")

	// Re-asigamos la variable fruit usando el operador de desreferenciacion en el puntero p - *p
	*p = "🥑"
	fmt.Printf("puntero: p\nTipo:\t\t\t %T \nValor:\t\t\t %v \nDireccion:\t\t %v \nDesreferen:\t\t %s \n", p, p, &p, *p)
	fmt.Println("_____________________________________________")

	// Imprimimos la variable  fruit para comprobar que se cambio su valor
	fmt.Printf("Variable: fruit\nTipo:\t\t\t %T \nValor:\t\t\t %s \nDireccion:\t\t %v \n", fruit, fruit, &fruit)
	fmt.Println("_____________________________________________")

}
