package main

import (
	"fmt"
	"strings"
)

func main() {
	res := sum(4, 6)
	fmt.Println("Total: ", res)

	min, may := convert("Jhon Leider g")
	fmt.Println("En minusculas: ", min)
	fmt.Println("En mayusculas: ", may)
}

// cuando 2 o mas parametros son del mismo tipo,
// solo se debe especificar en el ultimo parametro
func sum(num1, num2 int) int {
	return num1 + num2
}

// funcion con multiples retornos
func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}
