package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 5, 7))
}

// la funcion recibe cualquier cantidad de argumentos
// cada arguumento se añade a un slice
// se debe operar sabiendo que es un slice
func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}
