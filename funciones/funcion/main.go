package main

import "fmt"

// recibir funciones como parametros y retornar funciones

func main() {
	// Declaramos un slice de enteros
	nums := []int{1, 4, 70, 5, 67, 90, 2}

	//
	res := filter(nums, func(num int) bool {
		if num >= 10 {
			return true
		}

		return false
	})

	fmt.Println(res)

	// Una opcion para llamar una funcion que devuelve otra funcion es esta
	x := hello("Jhon Leider Guerrero")()
	fmt.Println(x)

	// Otra opcion es esta
	funcion := hello("Galia")
	fmt.Println(funcion())

	// Otra opcion es esta
	funcion1 := hello1("Galia")
	fmt.Println(funcion1("Guerrero"))

	// Otra opcion es esta
	funcion2 := hello1("Thor")("Guerrero")
	fmt.Println(funcion2)

}

// La funcion filter recibe un slice de enteros y una funcion callback,
// y devuelve un slice de enteros
// La funcion callback a su vez recibe un enerto  y devuelve un boolean
func filter(nums []int, callback func(int) bool) []int {
	// Se declara un slice de enteros
	result := []int{}

	// se hace un for range que itera sobre cada elemento de slice que llega por parametros
	for _, v := range nums {

		// Por cada valor del slice se llama a la funcion callback enviando el valor por parametro
		// la funcion callback devuelve true si el valor que se pasa es menor o igual a 10 y false si no
		if callback(v) {

			// Se agrega el valor si cumple la condicion de la funcion callback
			// usando la funcion append()
			result = append(result, v)
		}
	}

	// Para finalizar se regresa el slice de los numeros que cumplieron
	// la condicion de la funcion callback
	return result
}

// Funcion que retorna una funcion
// la funcion hello recibe un string y retorna una funcion
// que a su vez retorna un string
func hello(name string) func() string {
	return func() string {
		return "Hello " + name
	}
}

// igual que la anterior, pero la funcion que retorna recibe un parametro string
func hello1(name string) func(string) string {
	return func(text string) string {
		return "Hello " + name + " " + text
	}
}
